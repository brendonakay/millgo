package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"millgo/packages"
)

// Process CSV lines to struct
func stageOneChan(reader *csv.Reader, yc millgo.YamlConfig) <-chan millgo.AuditLog {
	// Channel sending data to initial transform stage
	stageOne := make(chan millgo.AuditLog)

	go func() {
		defer close(stageOne)
		for {
			line, err := reader.Read()
			if err == io.EOF {
				return
			} else if err != nil {
				log.Fatal(err)
			}
			auditLogStruct := millgo.AuditLog{
				EvidenceConstant:     millgo.EVIDENCE,
				AuditLogConstant:     millgo.AUDIT_LOG,
				Timestamp:    line[yc.AuditLog.Timestamp],
				PatientId:    line[yc.AuditLog.PatientId],
				EmployeeId:   line[yc.AuditLog.EmployeeId],
				AccessAction: line[yc.AuditLog.AccessAction],
			}
			stageOne <- auditLogStruct
		}
	}()
	return stageOne
}

// Transform the lines
func stageTwoChan(stageOneChan <-chan millgo.AuditLog, yc []byte) <-chan millgo.AuditLog {
	stageTwo := make(chan millgo.AuditLog)
	var clientRules = []millgo.Rule{}

	// TODO - Implement the rest of the field ops
	for k, v := range yc {
		switch k {
		case "constant":
			clientRules = append(clientRules, millgo.UseConstantRule{
				FieldName: v["field"],
				Constant:  v["value"],
				})
		case "date_fmt":
			clientRules = append(clientRules, millgo.ChangeDateFormatRule{
				FieldName: "Timestamp",
				NewDateFormat: "2006-01-02",
			})
		}
	}

	go func() {
		defer close(stageTwo)
		for line := range stageOneChan {
				//clientRules := []millgo.Rule{
				//	millgo.UseConstantRule{
				//		FieldName: "AccessAction", // TODO - Use YAML for these fields
				//		Constant:  "FOO",
				//	},
				//	millgo.ChangeDateFormatRule{
				//		FieldName: "Timestamp",
				//		NewDateFormat: "2006-01-02",
				//	},
				//}
			for _, rule := range clientRules {
				rule.Process(&line)
			}
			stageTwo <- line
		}
	}()
	return stageTwo
}

// Load to file
// TODO
//	- Eventually send to multiple different channels, e.g. CSV chan, Err chan...
//	- Buffered channel?
func stageThreeChan(stageTwoChan <-chan millgo.AuditLog) {
	outFile, err := os.OpenFile("test_output.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	w := csv.NewWriter(outFile)

	for line := range stageTwoChan {
		s := []string{
			line.EvidenceConstant,
			line.AuditLogConstant,
			line.Timestamp,
			line.PatientId,
			line.EmployeeId,
			line.AccessAction,
		}
		w.Write(s)
	}
	w.Flush()
}

func runAuditLogStages(reader *csv.Reader, yamlConfig millgo.YamlConfig) {
	yc := yamlConfig["audit_log"]

	// Get stageOne channel extract
	stageOne := stageOneChan(csvReader, yc)

	// Get stageTwo channel for transformation
	stageTwo := stageTwoChan(stageOne)

	// StageThree load
	stageThreeChan(stageTwo)
}

func main() {
	// Retrieve YAML file
	yamlBin, err := os.Open("data/example.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer yamlBin.Close()

	// Read YAML file
	yamlFile, err := ioutil.ReadAll(yamlBin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", yamlFile)

	// Parse YAML file to config struct
	yamlConfig := millgo.YamlConfig{}

	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range yamlConfig.FieldOps {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	// Open file
	fileHandle, err := os.Open("data/audit_log_example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	// Init CSV File reader
	csvReader := csv.NewReader(fileHandle)

	runAuditLogStages(csvReader, yamlConfig.FieldOps["audit_log"])
}

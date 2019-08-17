package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"millgo/packages"
)


// Line operations
func stageZeroChan(reader *csv.Reader, excludeList []string) <-chan []string {
	stageZero := make(chan []string)
	//errorChan := make(chan interface{})
	block := false

	go func() {
		defer close(stageZero)
		for {
			line, err := reader.Read()
			if err == io.EOF {
				return
			} else if err != nil {
				log.Fatal(err)
			}
			for _, i := range excludeList {
				if strings.Contains(strings.Join(line, ""), i) {
					block = true
					fmt.Println(line)
					// TODO: Get errorChan working!
					//errorChan <- line
					break
				}
			}
			if block != true {
				stageZero <- line
			} else {
				block = false
			}
		}
	}()
	return stageZero
}

// Unmarshall CSV line for field operations
func stageOneChan(stageZeroChan <-chan []string, yamlConfig millgo.YamlConfig) <-chan millgo.AuditLog {
	stageOne := make(chan millgo.AuditLog)
	yc := yamlConfig.AuditLog

	go func() {
		defer close(stageOne)
		for line := range stageZeroChan {
			auditLogStruct := millgo.AuditLog{
				EvidenceConstant: millgo.EVIDENCE,
				AuditLogConstant: millgo.AUDIT_LOG,
				Timestamp:        line[yc.Timestamp],
				PatientId:        line[yc.PatientId],
				EmployeeId:       line[yc.EmployeeId],
				AccessAction:     line[yc.AccessAction],
			}
			stageOne <- auditLogStruct
		}
	}()
	return stageOne
}

// Field operations / client Rules
func stageTwoChan(stageOneChan <-chan millgo.AuditLog, yamlConfig millgo.YamlConfig) <-chan millgo.AuditLog {
	stageTwo := make(chan millgo.AuditLog)
	yc := yamlConfig.FieldOps["audit_log"]
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
				FieldName:      v["field"],
				NewDateFormat:  v["new_date_format"],
				OrigDateFormat: v["orig_date_format"],
			})
		}
	}

	go func() {
		defer close(stageTwo)
		for line := range stageOneChan {
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

func runAuditLogStages(reader *csv.Reader, yamlConfig millgo.YamlConfig, excludeList []string) {
	stageZero := stageZeroChan(reader, excludeList)
	stageOne := stageOneChan(stageZero, yamlConfig)
	stageTwo := stageTwoChan(stageOne, yamlConfig)
	stageThreeChan(stageTwo)
	// Flush errorChan
	//for line := range err {
	//	fmt.Println(line)
	//}
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

	// Parse YAML file to config struct
	yamlConfig := millgo.YamlConfig{}

	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Open file
	fileHandle, err := os.Open("data/audit_log_example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	// TODO: Make this a file reader
	excludeList := []string{"foo", "bar", "Chart Review Tab"}

	// Init CSV File reader
	csvReader := csv.NewReader(fileHandle)

	runAuditLogStages(csvReader, yamlConfig, excludeList)
}

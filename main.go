package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	//"sync"

	"gopkg.in/yaml.v2"
	"millgo/packages"
)

// Process CSV lines to struct
func stageOneChan(reader *csv.Reader, yamlConfig millgo.YamlConfig) <-chan millgo.AuditLog {
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
				Evidence:     millgo.EVIDENCE,
				AuditLog:     millgo.AUDIT_LOG,
				Timestamp:    line[yamlConfig.AuditLog.Timestamp],
				PatientId:    line[yamlConfig.AuditLog.PatientId],
				EmployeeId:   line[yamlConfig.AuditLog.EmployeeId],
				AccessAction: line[yamlConfig.AuditLog.AccessAction],
			}
			stageOne <- auditLogStruct
		}
	}()
	return stageOne
}

// Transform the lines
// TODO
//	- Implement the rest of the fieldOps
func stageTwoChan(stageOneChan <-chan millgo.AuditLog) <-chan millgo.AuditLog {
	stageTwo := make(chan millgo.AuditLog)

	fieldOps := millgo.MyFieldOps()

	go func() {
		defer close(stageTwo)
		for line := range stageOneChan {
			for k, v := range fieldOps {
				switch k {
				case "useConstant":
					line.AccessAction = v.(func(string)string)("FOO")
				case "changeDateFormat":
					continue
				case "useMappedValue":
					continue
				case "removeFromValue":
					continue
				case "substituteInValue":
					continue
				}
			}

			stageTwo <- line
		}
	}()
	return stageTwo
}

// Load to file
// TODO
//	- Eventually send to multiple different areas.
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
			line.Evidence,
			line.AuditLog,
			line.Timestamp,
			line.PatientId,
			line.EmployeeId,
			line.AccessAction,
		}
		w.Write(s)
	}

	w.Flush()
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

	fmt.Printf("--- yaml:\n%v\n\n", yamlConfig)
	fmt.Printf("%v", yamlConfig.AuditLog.Timestamp)

	// Open file
	fileHandle, err := os.Open("data/audit_log_example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	// Init CSV File reader
	csvReader := csv.NewReader(fileHandle)

	// Get stageOne channel extract
	stageOne := stageOneChan(csvReader, yamlConfig)

	// Get stageTwo channel for transformation
	stageTwo := stageTwoChan(stageOne)

	stageThreeChan(stageTwo)

	// Read output from stageOne channel
	//for v := range stageTwo {
	//	fmt.Println(v)
	//}
	//fmt.Println("CHANNEL FLUSHED")
}

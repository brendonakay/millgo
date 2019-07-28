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
func stageOneChan(reader *csv.Reader, yamlConfig millgo.YamlConfig) (
	<-chan millgo.AuditLog, <-chan bool) {
	// Channel sending data to initial transform stage
	stageOne := make(chan millgo.AuditLog)
	endOfFile := make(chan bool)

	go func() {
		for {
			line, err := reader.Read()
			if err == io.EOF {
				fmt.Printf("END OF FILE")
				close(stageOne)
				break
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
		endOfFile <- true
		close(endOfFile)
	}()
	return stageOne, endOfFile
}

// Transform the lines
func stageTwoChan(stageOneChan <-chan millgo.AuditLog, endOfFile <-chan bool) (
	<-chan millgo.AuditLog, <-chan bool) {
	stageTwo := make(chan millgo.AuditLog)

	go func() {
		for {
			line := <-stageOneChan
			line.AccessAction = "TEST"

			stageTwo <- line
		}
	}()
	return stageTwo, endOfFile
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
	stageOne, eof := stageOneChan(csvReader, yamlConfig)

	// Get stageTwo channel for transformation
	stageTwo, eof := stageTwoChan(stageOne, eof)

	// Read output from stageOne channel
	for {
		select {
		case s := <-stageTwo:
			fmt.Println(s)
		case <-eof:
			return
		}
	}
}

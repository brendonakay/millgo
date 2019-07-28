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
// TODO: Can this be a goroutine? Why, yes it can. In fact, it might
//  work best as a generator function that returns a receive only chan.
// Perhaps implement a Select?
func stageOneChan(reader *csv.Reader, yamlConfig millgo.YamlConfig) (
	<-chan millgo.AuditLog, <-chan bool) {
	// Channel sending data to initial transform stage
	stageOne := make(chan millgo.AuditLog)
	endOfFile := make(chan bool)

	go func() {
		for {
			line, err := reader.Read()
			fmt.Printf("--- csv line:\n%s\n\n", line)
			if err == io.EOF {
				fmt.Printf("END OF FILE")
				close(stageOne)
				endOfFile <- true
				close(endOfFile)
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
	}()
	return stageOne, endOfFile
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

	// Get stageOne channel
	consumer, eof := stageOneChan(csvReader, yamlConfig)

	// Read output from stageOne channel
	for {
		select {
		case s := <-consumer:
			fmt.Println(s)
		case <-eof:
			return
		}
	}
}

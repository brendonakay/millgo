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

func main() {
	// Parse YAML
	yamlBin, err := os.Open("data/example.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer yamlBin.Close()

	yamlFile, err := ioutil.ReadAll(yamlBin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", yamlFile)

	//yaml := parseYaml(yamlFile)
	yamlConfig := millgo.YamlConfig{}
	//yamlConfig := make(map[interface{}]interface{})

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

	// Channel sending data to initial transform stage
	stageOne := make(chan millgo.AuditLog)
	endOfFile := make(chan bool)

	// Process CSV lines to struct
	// TODO: Can this be a goroutine? Why, yes it can. In fact, it might
	//  work best as a generator function that returns a receive only chan.
	// Perhaps implement a Select?
	for {
		line, err := csvReader.Read()
		fmt.Printf("--- csv line:\n%s\n\n", line)
		if err == io.EOF {
			fmt.Printf("END OF FILE")
			endOfFile <- true
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

	go func() {
		for{
			select {
			case s := <-stageOne:
				fmt.Println(s)
			case <-endOfFile:
				break
			}
		}
	}()
	// This will be the first transform stage. Receiving into Stage 1
	//for {
	//	fmt.Println(<-stageOne)
	//}
	//wg.Wait()
}

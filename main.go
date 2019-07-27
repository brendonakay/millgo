package main

import (
	//	"bufio"
	"encoding/csv"
	//"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"millgo/packages"
)

/*
// Unmarshal CSV line to millgo AccessLog struct
// return AccessLog struct CSV line
// TODO: rename?
func csvToAccessLogStruct(csvLine []string, yaml millgo.YamlConfig) millgo.AuditLog {
	csvLineStruct := millgo.AuditLog{
		Evidence:     millgo.EVIDENCE,
		AuditLog:     millgo.AUDIT_LOG,
		Timestamp:    yaml.Timestamp,
		PatientId:    yaml.PatientId,
		EmployeeId:   yaml.EmployeeId,
		AccessAction: yaml.AccessAction,
	}
	return csvLineStruct
}
*/

/*
// Parse YAML files.
// return millgo YamlConfig struct
func parseYaml(yamlFile []byte) millgo.YamlConfig {
	yamlConfig := make(map[interface{}]interface{})
	//yamlConfig := millgo.YamlConfig{}

	err := yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatal(err)
	}
	return yamlConfig
}
*/

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

	// Process CSV lines to struct
	// TODO: Can this be a goroutine?
	for {
		line, err := csvReader.Read()
		fmt.Printf("--- csv line:\n%s\n\n", line)
		if err == io.EOF {
			fmt.Printf("END OF FILE")
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
		// TODO: Print statement
		// This should probably send value to channel
		fmt.Printf("--- access log:\n%v\n\n", auditLogStruct)
	}
}

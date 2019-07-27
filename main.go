package main

import (
	//	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"millgo/packages"
)

func csvToAccessLogStruct(csvLine []string, yaml millgo.YamlConfig) millgo.AuditLog {
	csvLineStruct := millgo.AuditLog{
		Evidence:     EVIDENCE,
		AuditLog:     AUDIT_LOG,
		Timestamp:    yaml.Timestamp,
		PatientId:    yaml.PatientId,
		EmployeeId:   yaml.EmployeeId,
		AccessAction: yaml.AccessAction,
	}
	return csvLineStruct
}
func parseYaml(yamlFile []byte) millgo.YamlConfig {
	yamlConfig := millgo.YamlConfig{}

	err := yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatal(err)
	}
	return yamlConfig
}

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

	yaml := parseYaml(yamlFile)

	fmt.Printf("--- yaml:\n%v\n\n", yaml)

	// Open file
	fileHandle, err := os.Open("/Users/brendon/workspace/tmp/Jun-26-00_00_00-05_59_59-2019.psvlog")
	if err != nil {
		log.Fatalf(err)
	}
	defer fileHandle.Close()

	// Init CSV File reader
	csvReader, err := csv.NewReader(fileHandle)
	if err != nil {
		log.Fatalf(err)
	}
	defer fileHandle.Close()

	// Process CSV lines
	// TODO: Can this be a goroutine?
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		accessLogStruct := csvToAccessLogStruct(line, yaml)
		// TODO: Print statement
		fmt.Printf("--- access log:\n%v\n\n", accessLogStruct)
	}
}

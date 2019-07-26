package main

import (
	//	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"millgo/packages"
)

func parseYaml(yamlFile []byte) millgo.YamlConfig {
	yamlConfig := millgo.YamlConfig{}

	err := yaml.Unmarshal(yamlFile, &yamlConfig)
	fmt.Printf("--- yamlConfig:\n%v\n\n", yamlConfig)
	if err != nil {
		log.Fatal(err)
	}
	return yamlConfig
}

func main() {
	yamlBin, err := os.Open("example.yaml")
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

	/*
		// Open file
		fileHandle, err := os.Open("/Users/brendon/workspace/tmp/Jun-26-00_00_00-05_59_59-2019.psvlog")
		if err != nil {
			log.Fatalf(err)
		}
		defer fileHandle.Close()

		// Initialize scanner
		fileScanner := bufio.NewScanner(fileHandle)

		// Read out scanner
		for fileScanner.Scan() {
			fmt.Println(fileScanner.Text())
		}
	*/
}

package main

import (
	"bufio"
	"fmt"
    "log"
	"os"
)

func main() {

    // Open file
	fileHandle, err := os.Open("/Users/brendon/workspace/tmp/Jun-26-00_00_00-05_59_59-2019.psvlog")
    if err != nil {
        log.Fatal(err)
        defer fileHandle.Close()
    }

    // Initialize scanner
	fileScanner := bufio.NewScanner(fileHandle)

    // Read out scanner
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
}

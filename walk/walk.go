package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Walk function used for listing files recursively
	// No need to file info therefore file info ignored
	// Checked for error condition
	err := filepath.Walk(".",
		func(filePath string, _ os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("error occurred:", err)
			}
			// Filtered for only go files
			if len(filePath) > 4 && filePath[len(filePath)-3:] == ".go" {
				fmt.Println(filePath)

			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

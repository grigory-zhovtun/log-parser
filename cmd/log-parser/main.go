package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"log-parser/pkg/parser"
)

func main() {
	fmt.Println("Log Parser v.0.1")

	file, err := os.Open("access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	errorCount := 0
	successCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		_, err := parser.ParseLogLine(line)
		if err != nil {
			errorCount++
			continue
		}

		successCount++
		if successCount%10000 == 0 {
			fmt.Println(successCount)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success: %d, Errors: %d\n", successCount, errorCount)
}

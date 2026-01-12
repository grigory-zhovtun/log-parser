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

	stats := make(map[string]int)

	file, err := os.Open("access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	errorCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parser.ParseLogLine(line)
		if err != nil {
			errorCount++
			continue
		}
		stats[entry.IP]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success: %d, Errors: %d\n", len(stats), errorCount)
}

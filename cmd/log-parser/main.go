package main

import (
	"bufio"
	"fmt"
	"os"

	"log-parser/pkg/parser"
)

func main() {
	file, _ := os.Open("access.log")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	successCount := 0
	for scanner.Scan() {
		_, err := parser.ParseLogLine(scanner.Text())
		if err != nil {
			continue
		}
		successCount++
	}
	fmt.Printf("Parsed: %d\n", successCount)
}

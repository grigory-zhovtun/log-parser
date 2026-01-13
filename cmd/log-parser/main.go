package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"log-parser/pkg/parser"
)

type IPStat struct {
	IP    string
	Count int
}

func main() {
	fmt.Println("Log Parser v.0.1")

	filePath := flag.String("file", "access.log", "path to input file")
	limit := flag.Int("limit", 10, "max items to output")

	flag.Parse()

	stats := make(map[string]int)

	file, err := os.Open(*filePath)
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

	statsSlice := make([]IPStat, 0, len(stats))
	for ip, count := range stats {
		statsSlice = append(statsSlice, IPStat{IP: ip, Count: count})
	}

	sort.Slice(statsSlice, func(i, j int) bool {
		return statsSlice[i].Count > statsSlice[j].Count
	})

	fmt.Println("Top 10 IPs:")

	if len(statsSlice) < *limit {
		*limit = len(statsSlice)
	}

	for i := 0; i < *limit; i++ {
		fmt.Printf("%2d. %-15s %d requests\n", i+1, statsSlice[i].IP, statsSlice[i].Count)
	}
}

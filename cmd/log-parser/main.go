package logparser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Log Parser v.0.1")

	file, err := os.Open("access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++
		if count%10000 == 0 {
			fmt.Println(count)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("общее количество строк: %d\n", count)
}
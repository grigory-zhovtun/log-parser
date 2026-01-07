package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"flag"
)

func generator(numLines int, filePath string) (err error) {

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := 0; i < numLines; i++ {
        ip := getRandomIP()
        logTime := time.Now().Format(time.RFC3339)
        method := getRandomMethod()
        statusCode := getStatusCode()
        path := fmt.Sprintf("/api/v1/resource/%d", rand.Intn(1000))

        line := fmt.Sprintf("%s - [%s] \"%s %s HTTP/1.1\" %d\n", ip, logTime, method, path, statusCode)

        _, err := w.WriteString(line)
        if err != nil {
            return fmt.Errorf("ошибка записи в буфер: %w", err)
        }
    }

    return nil
}

func main() {
	var numLines int
	var filePath string

	flag.IntVar(&numLines, "n", 1000000, "number of lines to generate")
	flag.StringVar(&filePath, "file", "access.log", "path to output file")

	flag.Parse() 

	err := generator(numLines, filePath)
	if err != nil {
		fmt.Printf("ошибка: %v", err)
	}
}
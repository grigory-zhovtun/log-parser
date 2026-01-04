package generator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Generator() {

	file, err := os.Create("access.log")
	if err != nil {
		fmt.Printf("Ошибка создания файла: %v\n", err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := 0; i <= 100000; i ++ {
		ip := getRandomIP()
		log_time := time.Now().Format(time.RFC3339)
		method := getRandomMethod()
		status_code := getStatusCode()
		path := fmt.Sprintf("/api/v1/resource/%d", rand.Intn(1000))

        line := fmt.Sprintf("%s - [%s] \"%s %s HTTP/1.1\" %d\n", ip, log_time, method, path, status_code)

		w.WriteString(line)
    }
}
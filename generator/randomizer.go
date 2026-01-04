package generator

import (
	"fmt"
	"math/rand"
)

func getRandomNumber() int {
	roll := rand.Intn(256)

	return roll
}

func getRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", getRandomNumber(), getRandomNumber(), getRandomNumber(), getRandomNumber())
}

func getRandomMethod() string {
	http_methods := []string{"GET", "POST", "PUT", "DELETE"}

	return http_methods[rand.Intn(2)]
}

func getStatusCode() int {
	status_codes := []int{200, 201, 400, 404, 500}

	return status_codes[rand.Intn(len(status_codes))]
}
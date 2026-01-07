package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func init() {
	if runtime.Version() < "go1.20" {
		rand.Seed(time.Now().UnixNano())
	}
}

func getRandomNumber() int {
	roll := rand.Intn(256)

	return roll
}

func getRandomIP() string {
	numOne := strconv.Itoa(getRandomNumber())
	numTwo := strconv.Itoa(getRandomNumber())
	numThree := strconv.Itoa(getRandomNumber())
	numFour := strconv.Itoa(getRandomNumber())
	return fmt.Sprintf("%s.%s.%s.%s", 
		numOne, 
		numTwo, 
		numThree, 
		numFour,
	)
}

func getRandomMethod() string {
	httpMethods := []string{"GET", "POST", "PUT", "DELETE"}
	
	return httpMethods[rand.Intn(len(httpMethods))]
}

func getStatusCode() int {
	statusCodes := []int{200, 201, 400, 404, 500}

	return statusCodes[rand.Intn(len(statusCodes))]
}
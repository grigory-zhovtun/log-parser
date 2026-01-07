package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func funcInit() {
	if runtime.Version() < "go1.20" {
		rand.Seed(time.Now().UnixNano())
	}
}

func getRandomNumber() int {
	funcInit()

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
	funcInit()

	httpMethods := []string{"GET", "POST", "PUT", "DELETE"}
	
	return httpMethods[rand.Intn(len(httpMethods))]
}

func getStatusCode() int {
	funcInit()
	
	statusCodes := []int{200, 201, 400, 404, 500}

	return statusCodes[rand.Intn(len(statusCodes))]
}
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var httpMethods = []string{"GET", "POST", "PUT", "DELETE"}

func funcInit() {
	if runtime.Version() < "go1.20" {
		rand.Seed(time.Now().UnixNano())
	}
}

func getRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", 
		rand.Intn(256), 
		rand.Intn(256), 
		rand.Intn(256), 
		rand.Intn(256),
	)
}

func getRandomMethod() string {
	
	return httpMethods[rand.Intn(len(httpMethods))]
}

func getStatusCode() int {
	statusCodes := []int{200, 201, 400, 404, 500}

	return statusCodes[rand.Intn(len(statusCodes))]
}
package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	// benchmark code here

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Program executed in %v\n", executionTime)
}

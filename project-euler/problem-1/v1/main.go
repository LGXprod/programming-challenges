package main

import (
	"fmt"
	"time"
)

func main() {
	var executionTimeSum int64 = 0 // nanoseconds

	for i := 0; i < 5; i++ {
		startTime := time.Now()

		mo35Sum := getMo35Sum()

		endTime := time.Now()
		executionTimeSum += endTime.Sub(startTime).Nanoseconds()

		fmt.Print("Sum of Multiples of 3 and 5 under 1000 = ", mo35Sum, "\n\n")
	}

	meanExecutionTime := executionTimeSum / 5

	fmt.Printf("Program executed in %v nanoseconds (mean of 5 runs)\n", meanExecutionTime)
}

func getMo35Sum() int {
	mo35Sum := 0
	count := 2

	mo3 := 3
	mo5 := 5

	for mo3 < 1000 {
		if mo5 < 1000 {
			mo35Sum += mo5
			mo5 = 5 * count
		}

		mo35Sum += mo3
		mo3 = 3 * count

		count += 1
	}

	return mo35Sum
}

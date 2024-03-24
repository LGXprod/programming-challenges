package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(getRandomNums(-10, 10, 10))
	}
}

func getRandomNums(lowerBound int, upperBound int, length int) []int {
	var randomNums []int

	negativeOffset := 0

	if lowerBound < 0 {
		negativeOffset = -1 * lowerBound
	}

	for i := 0; i < length; i++ {
		randomNum := lowerBound + rand.Intn(upperBound+1+negativeOffset)
		randomNums = append(randomNums, randomNum)
	}

	return randomNums
}

package services

import (
	"fmt"
	"time"
)

// Fibonacci function calculate a n position of Fibonacci series
func Fibonacci(n float64) float64 {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// TimeTrack measure the execution time of a function
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

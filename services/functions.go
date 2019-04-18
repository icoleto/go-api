package services

// Fibonacci function calculate a n position of Fibonacci series
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

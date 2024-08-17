package main

import "math"

type Condition func(n int) bool

// Function to check if a number is even
func isEven(number int) bool {
	return number%2 == 0
}

// Function to check if a number is odd
func isOdd(number int) bool {
	return !isEven(number)
}

// Function to check if a number is prime
func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(number)))
	for i := 3; i <= sqrtN; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is a multiple of n
func isMultipleOfN(number, n int) bool {
	return number%n == 0
}

// Function to check if a number is greater than n
func isGreaterThanN(number, n int) bool {
	return number > n
}

// Funtions to filter a list of integers based on all the conditions provided
func filterAll(numbers []int, conditions ...Condition) []int {
	var outputs []int
	for _, number := range numbers {
		var matchesCondition = true
		for _, condition := range conditions {
			if !condition(number) {
				matchesCondition = false
				break
			}
		}
		if matchesCondition {
			outputs = append(outputs, number)
		}
	}

	return outputs
}

package main

import (
	"fmt"
)

// https://one2n.io/go-bootcamp/go-projects/basic-number-filtering/basic-number-filtering-exercise

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter out only even numbers
	conditions := []Condition{isEven}
	evenNumbers := filterAll(numbers, conditions...)
	fmt.Println(evenNumbers)

	// Filter out only odd numbers
	conditions = []Condition{isOdd}
	oddNumbers := filterAll(numbers, conditions...)
	fmt.Println(oddNumbers)

	// Filter out only prime numbers
	conditions = []Condition{isPrime}
	primeNumbers := filterAll(numbers, conditions...)
	fmt.Println(primeNumbers)

	// Filter out only odd prime numbers
	conditions = []Condition{isOdd, isPrime}
	oddPrimeNumbers := filterAll(numbers, conditions...)
	fmt.Println(oddPrimeNumbers)

	// Filter out only even multiples of 5
	isMultipleOfFive := func(n int) bool { return isMultipleOfN(n, 3) }
	conditions = []Condition{isEven, isMultipleOfFive}

	evenMultiplesOfFive := filterAll(numbers, conditions...)
	fmt.Println(evenMultiplesOfFive)

	// Filter out only odd multiples of 3 greater than 10
	isGreaterThanTen := func(n int) bool { return isGreaterThanN(n, 10) }
	isMultipleOfThree := func(n int) bool { return isMultipleOfN(n, 3) }
	conditions = []Condition{isOdd, isMultipleOfThree, isGreaterThanTen}

	oddMultiplesOfThreeGreaterThanTen := filterAll(numbers, conditions...)
	fmt.Println(oddMultiplesOfThreeGreaterThanTen)
}

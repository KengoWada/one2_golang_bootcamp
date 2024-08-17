package main

import (
	"reflect"
	"testing"
)

func TestIsEven(t *testing.T) {
	got := isEven(2)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = isEven(3)
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsOdd(t *testing.T) {
	got := isOdd(4)
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = isOdd(5)
	want = true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsPrime(t *testing.T) {
	got := isPrime(2)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = isPrime(4)
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsMultipleOfN(t *testing.T) {
	got := isMultipleOfN(2, 2)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = isMultipleOfN(3, 2)
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsGreaterThanN(t *testing.T) {
	got := isGreaterThanN(2, 2)
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = isGreaterThanN(3, 2)
	want = true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFilterAll(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Filter all odd prime numbers
	conditions := []Condition{isOdd, isPrime}
	want := []int{3, 5, 7, 11, 13, 17, 19}
	got := filterAll(numbers, conditions...)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	isGreaterThanTen := func(n int) bool { return isGreaterThanN(n, 10) }
	conditions = []Condition{isOdd, isGreaterThanTen}
	want = []int{11, 13, 15, 17, 19}
	got = filterAll(numbers, conditions...)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

package main

import (
	"testing"
)

func TestVerifyPlaygroundPermissions(t *testing.T) {
	p := verifyPlaygroundPermissions(3, 1.2)
	result := []string{"carrossel", "montanha-russa"}
	if len(p) != len(result) {
		t.Errorf("result %v, but got %v", result, p)
	}

	for i := range p {
		if p[i] != result[i] {
			t.Errorf("result %v, but got %v", result, p)
		}
	}
}

func TestFactorialNumber(t *testing.T) {
	num := 5
	factorial := factorialNumber(num)
	result := 120
	if factorial != result {
		t.Errorf("Factorial of %v = %v, but got %v", num, result, factorial)
	}
}

func TestIsPrimeNumber(t *testing.T) {
	num := 6
	isprimenumber := isPrimeNumber(num)
	result := false
	if isprimenumber != result {
		t.Errorf("%v is not a prime number, but got %v", num, result)
	}
}

func TestHasSequence(t *testing.T) {
	sequence := []int{1, 2, 3, 4}
	hassequence := hasSequence(sequence)
	result := false
	if hassequence != result {
		t.Errorf("%v has not a sequence, but got %v", sequence, hassequence)
	}
}

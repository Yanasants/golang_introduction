package main

import "fmt"

func main() {
	// verifyPlaygroundPemissions
	age := 3
	height := 1.2
	playground := verifyPlaygroundPermissions(age, height)
	fmt.Printf("Idade: %v, Altura: %v", age, height)
	fmt.Printf("\nPermissão para brincar em: %v", playground)
	// factorialNumber
	factorial := factorialNumber(3)
	fmt.Println("\nFatorial: ", factorial)
	// isPrimeNumber
	dividers := isPrimeNumber(3)
	fmt.Println("\nÉ primo: ", dividers)
	// hasSequence
	sequence := []int{3, 2, 5, 5}
	hassequence := hasSequence(sequence)
	fmt.Println("\nPossui nº repetido em sequência: ", hassequence)

}

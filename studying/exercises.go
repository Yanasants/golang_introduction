package main

type playground_list []string

func verifyPlaygroundPermissions(num_age int, height float64) playground_list {

	playground := playground_list{}

	if height >= 1 {
		if num_age >= 3 {
			playground = append(playground, "carrossel")
		}
	}
	if height <= 1.3 {
		if num_age >= 4 {
			if num_age <= 7 {
				playground = append(playground, "piscina de bolinha")
			}
		}
	}
	if height >= 1.1 {
		playground = append(playground, "montanha-russa")
	}
	return playground
}

type numbers_list []int

func factorialNumber(num int) int {
	numbers := numbers_list{}
	for i := num; i > 0; i-- {
		numbers = append(numbers, i)
	}

	factorial := 1
	for _, n := range numbers {
		factorial *= n
	}
	return factorial
}

func isPrimeNumber(num int) bool {
	dividers := numbers_list{}
	for i := num; i > 0; i-- {
		if num%i == 0 {
			dividers = append(dividers, i)
		}
	}
	if len(dividers) == 2 {
		return true
	} else {
		return false
	}
}

func hasSequence(sequence numbers_list) bool {
	for i := 1; i < len(sequence); i++ {
		if sequence[i] == sequence[i-1] {
			return true
		}
	}
	return false
}

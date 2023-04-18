package main

type numbers_list []int

func createNumberList(min int, max int) numbers_list {
	numbers := numbers_list{}
	for i := min; i <= max; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

package main

import "fmt"

func main() {
	numList := createNumberList(0, 15)
	for _, i := range numList {
		if i%2 == 0 {
			fmt.Println(i, "is even.")
		} else {
			fmt.Println(i, "is odd.")
		}
	}
}

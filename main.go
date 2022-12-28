package main

import (
	"fmt"
)

// takes slice of integers and bubble sorts them
func BubbleSort(input []int) []int {
	n := len(input)
	swapped := true

	for swapped {
		swapped = false

		// walk through input slice
		for i := 0; i < n-1; i++ {

			// if element i is bigger than element i+1 swap them
			if input[i] > input[i+1] {

				fmt.Println("Bubble Up!", input[i])

				input[i], input[i+1] = input[i+1], input[i]

				swapped = true
			}
		}
	}
	return input
}

func main() {
	fmt.Println("Go Bubble Sort")

	unsortedInput := []int{8, 2, 9, 3, 1, 6, 7}
	fmt.Println("Input: ", unsortedInput)

	sortedOutput := BubbleSort(unsortedInput)
	fmt.Println("Output: ", sortedOutput)
}

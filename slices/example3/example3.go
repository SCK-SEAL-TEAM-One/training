package main

import "fmt"

func main() {

	// Create a slice with initial values
	numbers := []int{10,15,5,9,2,4,8,0,11}
	fmt.Printf("numbers :%+v\n", numbers)

	// Take a slice of numbers. Start from index 5.
	// Parameters are [starting_index :]
	firstSet := numbers[5:]
	fmt.Printf("numbers :%+v\n", firstSet)

	// Take a slice of numbers. We want just indexes 2 and 5.
	// Parameters are [starting_index : (starting_index + length)]
	secondSet := numbers[2:5]
	fmt.Printf("numbers :%+v\n", secondSet)

	// Take a slice of numbers. Start from index 0 to index 4.
	// Parameters are [:end_index]
	thirdSet := numbers[:4]
	fmt.Printf("numbers :%+v\n", thirdSet)

}
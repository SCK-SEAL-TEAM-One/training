// Sample program to show how anonymous functions and closures work.

package main

import "fmt"

func main() {
	var n int

	// Declare an anonymous function and call it.
	func() {
		fmt.Println("Direct:", n)
	}()

	n++
	// Declare an anonymous function and assign it to a variable.
	f := func(x int) {
		fmt.Println("Variable:", x)
	}

	f(5)
}

// Sample program to show how to declare methods against
// a named type.
package main

import "fmt"

// human is a struct to bind methods to.
type human struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d human) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *human) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {

	// Declare a variable of type human.
	d := human{
		name: "Bill",
	}
	c := human{
		name: "lek",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the Compiler is Doing:")

	// This is what Go is doing underneath.
	human.displayName(c)
	(*human).setAge(&c, 24)

	// =========================================================================

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	f1()

	// =========================================================================

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	d.name = "Sammy"

	// Call the method via the variable. We see the change.
	f2(45)
}

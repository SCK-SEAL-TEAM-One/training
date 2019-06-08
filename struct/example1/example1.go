package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type student struct {
	name   string
	code   int
	gpa    float64
	status bool
}

func main() {
	// Declare a variable of type example set to its
	// zero value.
	var s1 student

	// Display the value.
	fmt.Printf("Student %+v\n", s1)

	// Initialize a struct by supplying the value of all the struct fields.
	var p = Person{"Rajeev", "Singh", 26}

	// Display the value.
	fmt.Printf("Person %+v\n", p)

	// Declare a variable of type example and init using
	// a struct literal.
	s2 := student{
		name:   "Nareenart",
		code:   1044,
		gpa:    4.00,
		status: true,
	}

	// Display the field values.
	fmt.Println("Name", s2.name)
	fmt.Println("Code", s2.code)
	fmt.Println("GPA", s2.gpa)
	fmt.Println("Status", s2.status)

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	s3 := struct {
		name   string
		code   int
		gpa    float64
		status bool
	}{
		name:   "Panumars",
		code:   1002,
		gpa:    3.59,
		status: true,
	}

	// Display the values.
	fmt.Printf("%+v\n", s3)
	fmt.Println("Name", s3.name)
	fmt.Println("Code", s3.code)
	fmt.Println("GPA", s3.gpa)
	fmt.Println("Status", s3.status)
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Declare variables that are set to their zero value.
	var amount int
	var name string
	var price float64
	var status bool

	fmt.Printf("var amount int \t %T [%v], size: %d byte\n", amount, amount, unsafe.Sizeof(amount))
	fmt.Printf("var name string \t %T [%v], size: %d byte\n", name, name, unsafe.Sizeof(name))
	fmt.Printf("var price float64 \t %T [%v], size: %d byte\n", price, price, unsafe.Sizeof(price))
	fmt.Printf("var status bool \t %T [%v]\n\n", status, status)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	money := 10
	say := "Hello"
	gpa := 3.99
	active := true

	fmt.Printf("money := 10 \t %T [%v]\n", money, money)
	fmt.Printf("say := \"hello\" \t %T [%v]\n", say, say)
	fmt.Printf("gpa := 3.14159 \t %T [%v]\n", gpa, gpa)
	fmt.Printf("active := true \t %T [%v]\n\n", active, active)

	// Specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

// อ้างอิง https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe

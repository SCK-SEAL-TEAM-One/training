package main

import (
	"encoding/json"
	"fmt"
)

type example1 struct {
	flag    bool
	counter int16
	pi      float32
}

type Example2 struct {
	Flag    bool
	Counter int16
	Pi      float32
}

type example3 struct {
	Flag    bool    `json:"flag"`
	Counter int16   `json:"counter"`
	Pi      float32 `json:"pi"`
}

func main() {
	// Declare a variable of type example and init using
	// a struct literal.
	e1 := example1{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Call outside package encoding/json
	e1Json, err := json.Marshal(e1)

	// Check error after call package encoding/json
	if err != nil {
		fmt.Println("json.Marshal: ", err)
	}

	// Display the values.
	fmt.Printf("e1Json: %+v\n", string(e1Json))

	// =========================================================================

	// Declare a variable of type example and init using
	// a struct literal.
	e2 := Example2{
		Flag:    true,
		Counter: 10,
		Pi:      3.141592,
	}

	// Call outside package encoding/json
	e2Json, err := json.Marshal(e2)

	// Check error after call package encoding/json
	if err != nil {
		fmt.Println("json.Marshal: ", err)
	}

	// Display the values.
	fmt.Printf("e2Json: %+v\n", string(e2Json))

	// =========================================================================

	// Declare a variable of type example and init using
	// a struct literal.
	e3 := example3{
		Flag:    true,
		Counter: 10,
		Pi:      3.141592,
	}

	// Call outside package encoding/json
	e3Json, err := json.Marshal(e3)

	// Check error after call package encoding/json
	if err != nil {
		fmt.Println("json.Marshal: ", err)
	}

	// Display the values.
	fmt.Printf("e3Json: %+v\n", string(e3Json))
}

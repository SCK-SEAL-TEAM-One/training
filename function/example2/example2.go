package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	id   int
	name string
}

func main() {

	// getUser Don't care about the user value.
	_, err := getUser()
	if err != nil {
		return
	}

	fmt.Println("err is: ", err)

	// getUser Don't care about the error value.
	u, _ := getUser()

	fmt.Printf("user is: %+v\n", u)
}

// getUser returns a pointer of type user.
func getUser() (*user, error) {
	return &user{1432, "Betty"}, nil
}

package main

import "fmt"

func main() {
	age := 22

	// for now go does not have ternary operator for the sake of simplicity
	if age >= 18 {
		fmt.Println("person is adult")
	} else if age < 18 && age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is chid")
	}

	// we can declare the varible in the if statement and then directly use then
	if age := 15 ; age >= 18 {
		fmt.Println("person is adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is child")
	}

	var role string = "admin"
	var hasPermission bool = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}
}
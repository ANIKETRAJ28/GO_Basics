package main

import "fmt"

func main() {
	// string
	// long hand
	var name string = "golang"
	fmt.Println(name)
	// type infer
	var name1 = "golang"
	fmt.Println(name1)
	// short hand
	name2 := "golang"
	fmt.Println(name2)

	// int
	var age int = 32
	fmt.Println(age)
	age2 := 32
	fmt.Println(age2)

	// why long hand?
	// for declaration we need to use long hand
	var age3 int
	age3 = 32
	fmt.Println(age3)

	// float
	price := 50.5
	fmt.Println(price)
}
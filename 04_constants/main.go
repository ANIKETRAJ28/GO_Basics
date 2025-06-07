package main

import "fmt"

func main() {
	const name string = "golang"
	// name = "javascript" // not allowed
	const (
		PORT = 5000
		host = "localhost"
	)

	fmt.Println(PORT, host)
}
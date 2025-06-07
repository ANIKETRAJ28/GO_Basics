package main

import "fmt"

// only for loop is present in go
// We use for loop as while loop also

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic loop
	for i := 0 ; i <= 10 ; i++ {
		if i == 5 {
			continue
		}
		if i == 10 {
			break
		}
		fmt.Println(i)
	}

	// range loop
	// [0, 3)
	for i := range 3 {
		fmt.Println(i)
	}
}
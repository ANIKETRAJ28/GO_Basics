package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	// simple switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: 
		fmt.Println("other")
	}
	// multiple condition switch
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday: 
			fmt.Println("weekend")
		default:
			fmt.Println("workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		// i.(type) -> gives the type of the variable 
		// which is only used inside switch statement
		switch i.(type) {
			case int:
				fmt.Println("int data type")
			case string:
				fmt.Println("string data type")
			case bool:
				fmt.Println("bool data type")
			default:
				fmt.Println(i, " has other data type")
		}
	}

	whoAmI("golang")
	whoAmI(3.24)
}
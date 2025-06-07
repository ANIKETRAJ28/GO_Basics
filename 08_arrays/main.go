package main

import "fmt"

func main() {

	// array usecase: 
	// the size or length is known
	// memory optimiation
	// constant time access

	// by default falsy values are initialized in arrays
	// for case of int -> 0
	var nums[4] int
	fmt.Println(len(nums))
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)
	// for case of bool -> false
	var vals[4] bool
	fmt.Println(vals)
	// for case of string -> ""
	var strs[4] string
	fmt.Println(strs)
	// initialize the array
	arr := [3] int {1,2,3}
	fmt.Println(arr)
	brr := arr // assign array to another variable
	fmt.Println(brr)
}
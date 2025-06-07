package main

import "fmt"

func main() {
	// slice -> dynamic array
	// most used construct in go
	// useful methods

	var nums []int
	// uninitiazed slice is by default nil (null in different languages)
	fmt.Println(nums == nil)
	fmt.Println(nums)
	fmt.Println(len(nums))

	// initialize the size for slice
	// make is the inbuilt function used to initialize the slice with default size, 
	// the second argument is the size and by default falsy values are set
	var arr = make([]int, 2)
	fmt.Println(arr)

	// the third is capacity, i.e what is the initial capacity of the slice
	var brr = make([]int, 2, 4)
	brr = append(brr, 1) // 0 0 1 -> 4
	fmt.Println(brr, cap(brr))
	brr = append(brr, 2) // 0 0 1 2 -> 4
	fmt.Println(brr, cap(brr))
	brr = append(brr, 3) // 0 0 1 2 3 -> 8
	fmt.Println(brr, cap(brr))
	brr = append(brr, 4) // 0 0 1 2 3 4 -> 8
	fmt.Println(brr, cap(brr))
	brr = append(brr, 5) // 0 0 1 2 3 4 5 -> 8
	fmt.Println(brr, cap(brr))

	var crr = make([]int, 0, 5)
	crr = append(crr, 1)
	crr = append(crr, 2)
	crr = append(crr, 3)
	crr = append(crr, 4)
	crr = append(crr, 5)
	fmt.Println(crr, cap(crr), len(crr))
	crr = append(crr, 6)
	crr = append(crr, 7)
	crr = append(crr, 8)
	fmt.Println(crr, cap(crr), len(crr))
}
package main

import "fmt"

func main() {
	// Array: fixed size of 3 integers
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slices: dynamic size
	slice := []int{4, 5, 6}
	fmt.Println("Slice:", slice)

	// Append to slice
	slice = append(slice, 7, 8)
	fmt.Println("Slice after append:", slice)

	// Slicing a slice
	subSlice := slice[1:4]
	fmt.Println("Sub-slice:", subSlice)

	// Length and capacity
	fmt.Println("Length:", len(slice))   // no. of element
	fmt.Println("Capacity:", cap(slice)) // Double from beginning
}

package main

import "fmt"

// NOTE:
/**
Array - fixed size sequence of elements of same type.
Slice - dynamic size sequence of elements of same type.
*/

func main() {
	// array

	var a [5]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	// slices in go.
	var s []int
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)

}

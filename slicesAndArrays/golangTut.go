package main

import "fmt"

/**
NOTE: Arrays are fixed size collection of elements.
You can't mix int and string. Will be of one type only.
*/

func arrayExamples() {
	// decalartoin 1
	var a [4]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	// short hand
	b := [3]int{1, 2, 3}
	fmt.Println(b)

	// optimized one. Compiler will figure out the length itself.
	c := [...]int{10, 20, 30, 40}
	fmt.Println(c)

	// NOTE: arrays are value types not referenced types. When they are assigned to a new array,
	// a copy of original array is assigned to the new varibale.
	// changes to the new variable, wont be reflected in the original array.
	// NOTE: same holds true when passing in a function as well.
	d := c
	d[0] = 100
	fmt.Println("c is", c)
	fmt.Println("d is", d)

	// Looping over arrays
	// way 1
	for i := 0; i < len(a); i++ {
		fmt.Println("a is", a[i])
	}

	// you will get both index, value
	for index, val := range a {
		fmt.Printf("index: %d, val: %d\n", index, val)
	}
}

func sliceExamples() {
	// NOTE:
	// slice is basically a wrapper over existing arrays.
	// they are just references to exiting arrays.

	// an array
	a := [5]int{76, 77, 78}

	// slice
	// [0,1]
	var b []int = a[0:2]
	fmt.Println(b)

	// NOTE: way 2 of creatina a slice
	c := []int{2, 3, 10}
	fmt.Println(c)

	// NOTE: a slice doesn't owns data of its own.
	//  ANy changes made to a slice will reflect on the original array
	darr := [...]int{10, 12, 13, 14, 15, 16, 16}
	dslice := darr[0:5]
	fmt.Println("array before", darr)

	for ind, _ := range dslice {
		dslice[ind]++
	}

	fmt.Println("array after", darr)

	// Example 2
	numa := [5]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which
	nums2 := numa[:]

	fmt.Println("array before change", numa)
	nums1[0] = 100
	fmt.Println("array after modificatin to array1", nums1)
	nums2[1] = 200
	fmt.Println("array after modificatin to array2", nums2)

	// FIXME: most important
	// length and capacity of a slice
	// length - is the number of current elements
	// capacity - is the size of the original array
	fmt.Printf("Len of nums1: %d, Cap of nums1:%d", len(nums1), cap(nums2))

	// NOTE: create your own slice
	// (type, len, cap)
	// cap is optional and defaults to length.
	sliceArray := make([]int, 5)
	fmt.Println()

	// NOTE: append will append at 6 element
	// first 5 will be 0
	sliceArray = append(sliceArray, 10)
	sliceArray = append(sliceArray, 20)

	// NOTE: during append under the hood a new array is created.
}

func main() {
	//arrayExamples()
	sliceExamples()
}

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

func main() {
	arrayExamples()
}

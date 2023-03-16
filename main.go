package main

import "fmt"

//fmt stands for format

func main() {

	//	type IDnum int
	//var pid IDnum

	//x := 10 initialization and declaration at the same time.
	// var temp string

	// var x int = 1
	// var y int

	// var ip *int
	// ip = &x
	// y = *ip

	// fmt.Printf(*y)

	// // new() - creates a variable and returns a pointer to that variable
	// ptr = new(int)
	// *ptr = 3

	// fmt.Printf("%d", temp)
	fmt.Printf("Hello, world!\n")

	/*
		Garbage collection in Python, JVM works in the way that it keeps track of the references/ptrs
		to a variable, when there are no more references to it, then the GC deacllocates it
	*/

	/* In Go, GC is done in the Background*/

	fmt.Printf("Hello, world %s", "Rahul")
	x := "Rahul"

	if x == "VERMA" {
		fmt.Printf("OK")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	var appleNum int
	fmt.Printf("Enter the number of apples")
	// num is number of scanned items and err is the error code
	num, err := fmt.Scan(&appleNum)
	fmt.Printf("%d", appleNum)
	fmt.Printf("%s", err)
	fmt.Printf("%d", num)

	fmt.Printf("*******WEEK3******")
	fmt.Printf("***ARRAYS")

	// var x[5] int
	// x[0]=1

	// temp2 := [3]int {1, 2, 3}
	// for i,v range temp2 {
	// 	fmt.Printf("%d %d", i,v)
	// }

	arr := [5]string{"a", "b", "c", "d", "ed"}
	s1 := arr[1:3]
	fmt.Printf("%s", s1)
	fmt.Printf("%d %d", len(s1), cap(s1))

}

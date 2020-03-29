//package main

import (
	"fmt"
	"math/rand"
)

func foo() {
	//NOTE - P, S is capitalized. Capitalization is sort of considered an external
	//thing
	//fmt.Println("The squareroot for 4 is", math.Sqrt(4)
}

/*
Topics
Syntax
TYpe
Pointers 
*/

// or fund add(x,y float64)
func add(x float64, y float64) float64 {
	return x + y
}

// note you can return two things in Go.
func multiple(x, y string) (string, string) {
	return x, y
}

//func main() {

	// you can type case
	var a int = 62
	var b float64 = float64(a)

	var num1, num2 float64 = 5.6, 9.5

	//or but you can't change it's type later.
	num3 := 15.6

	word1, word2 := "rahul", "verma"
	fmt.Println(multiple(word1, word2))

	fmt.Println("Addition of two float", add(num1, num2))
	// it will return 81 everytime. You need to provide a seed
	fmt.Println("generate a random number", rand.Intn(100), num3)
	foo()
}

// In python it is main()
// but in go main will always exist. Execution will start from main.

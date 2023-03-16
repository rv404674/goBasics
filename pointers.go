package main

import "fmt"

func main() {
	var x int = 10
	var ptr *int = &x

	fmt.Println("memory address", ptr)
	fmt.Println(*ptr)

}

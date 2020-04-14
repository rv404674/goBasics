package main

import "fmt"

//func main() {
	// doesn't have a while loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	//use it like while
	x := 10
	for {
		fmt.Println("Do stuff", x)
		x += 3
		if x > 25 {
			break
		}

	}

	for x = 5; x < 25; x += 5 {
		fmt.Println("dfsd")
	}

}

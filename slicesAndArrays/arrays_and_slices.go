package main

import "fmt"

func changeLocal(a [3]int) {
	a[0] = 100
}

func main() {
	// type 1 of declaring
	var a [3]int
	fmt.Println(a) // will o/p [0,0,0]

	// short hand declaration
	b := [3]int{12, 34, 45}
	// or b := [12]int
	fmt.Println(b)

	// NOTE - arrays are value types in go, see below
	c := [3]int{12, 34, 45}
	d := c

	d[0] = 24
	// This won't reflect the original array i.e c, unlike ptyhon
	fmt.Println(c[0], d[0])

	// Similarly arrays are passed by value to functions
	fmt.Println("Array before calling funtion", c)
	changeLocal(c)
	fmt.Println("Array after calling function", c)

	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	// NOTE important - Use Range to iterate over array
	for i, j := range c {
		fmt.Println(i, j)
	}

	// or if you want to ignore index
	for _, val := range c {
		fmt.Println(val)
	}

	// multidimensinal array
	tempA := [3][3]string{{"abc", "xyz"}, {"pqr", "xyz"}, {"dfs", "dfs"}}
	fmt.Println(tempA)

}

package main

import "fmt"

/**
NOTE: an interface can hold any value of any type.
fmt.Println() uses interface
NOTE: you can use any instead of interfac{}
*/

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	var i interface{}
	i = 42
	i = "hello"
	i = true
	fmt.Println(i)

	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}

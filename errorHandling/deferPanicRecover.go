package main

import (
	"errors"
	"fmt"
)

/**
RESOURCES - https://blog.logrocket.com/error-handling-golang-best-practices/
*/

// NOTE:
// go's ideology is that adding exceptions like try/catch/finally would
// result in complex code and encourage programmers to label too many basic
// error, like failiing to open a file as exceptional.

// you should not use defer/panic/recover as try/catch/finally.
// rather use it FIXME: very judiciously.

// NOTE:
// DEFER is using for cleaning up resources.
// 1. Defer is a lang mechanism that puts your function into callstack.
// 2. Each defer is executed in reverse order when the host function finihes.

func deferDemo() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	fmt.Println("Main Host Function")
}

// NOTE: PANIC
// panic is a builtin function that stops the normal flow execution.
// Use it in rare scenarios when nothing can be done.
// NOTE: the flow is first all the defers will be executed then panic will be called

/**
Panics are something I very rarely use within my day-to-day
Go application development. Most of the time,
catching an error from whatever I’m calling and
then logging it out and possibly emitting a metric that lets me know how often this is
happening is the preferential option.
*/
func panicDemo() {
	fmt.Println("Panic Demo")
	defer fmt.Println("closing cache")
	panic(errors.New("Panic Demo"))
}

// NOTE: Recover
// it is used to recover from panic. Basically the leftover statements can execute then
// EX
func sillySusan() {
	defer func() {
		if r := recover(); r != nil {
			// panic occured
			fmt.Println("Susan recovered from the panic")
			fmt.Println("silly susan handles the panic gracefully")
		}
	}()

	fmt.Println("silly susan called")
	panickingPeter()
	fmt.Println("silly susan finished")
}

func panickingPeter() {
	fmt.Println("panicking peter called")
	panic("Oh no")
	fmt.Println("panicking peter finished")
}

func main() {
	fmt.Println("cascading panics")
	sillySusan()
	//panicDemo()
	//deferDemo()

	//fmt.Println("Saving the recover thing")
	//defer handlePanic()
	//defer deferDemo()

	//panicDemo()
}

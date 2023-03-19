package main

import (
	"errors"
	"fmt"
)

/**
RESOURCES
1. https://hussachai.medium.com/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76
*/

// in go there is an error interface
type error interface {
	Error() string
}

// NOTE: Creating your own error

type MyError struct {
	Msg string
}

// create a function and associate it with the above
func (error *MyError) Error() string {
	return error.Msg
}

// Now you can construct an erorr object using MyError struct
func ThisFunctionReturnError() error {
	return &MyError{Msg: "This is a runtime error"}
}

// NOTE: easiest
func ThisRuntimeError() error {
	return errors.New("This is a runtime error")
}

// THe usual way is returning two things from a function.
func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, fmt.Errorf("Divisible by zero is not possible")
	}

	return x / y, nil
}

// NOTE: this is how you define a custome rror - Part2
type CustomError struct {
	Code   string
	Msg    string
	Status int
}

// create a function and associate it.
func (error *CustomError) Error() string {
	return error.Msg
}

func CustomerErrorInstance(msg string) *CustomError {
	return &CustomError{
		Code:   "rbe-001",
		Msg:    msg,
		Status: 400,
	}
}

func main() {
	ans, err := divide(5, 0)
	if err != nil {
		fmt.Println(err.Error(), ans)
	}

	customerErr1 := CustomerErrorInstance("Pool size exhausted")
	fmt.Println(customerErr1.Code)
	fmt.Println(customerErr1)

	customerErr2 := CustomerErrorInstance("Runtime Error")
	fmt.Println(customerErr2)

}

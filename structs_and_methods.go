package main

import "fmt"

/*
In go there are two type of methods.
1. Value Recievers - They can just recieve values and do calculations on it, they
CANT MODIFY ANY VALUE
2. Pointer Recievers - If you want to change a values in the struct.
*/

type student struct {
	name     string
	rollno   int
	dept     string
	ds_marks int
	os_marks int
}

func (s student) totalmarks() int {
	return (s.ds_marks + s.os_marks)
}

func main() {

	student_one := student{
		name:     "Rahul Verma",
		rollno:   151364,
		dept:     "Computer Science",
		os_marks: 99,
		ds_marks: 100}

	//student_two := student{"Sumedha Sagar", 151020, "ECE"}

	fmt.Printf("%s roll number is %d", student_one.name, student_one.rollno)
	fmt.Println(student_one.rollno)
	fmt.Println(student_one.name)

	fmt.Println("MARKS OF STUDENT1", student_one.totalmarks())
}

package main

import "fmt"

//Pointer reciever is used to modify any existing value.

type Student struct {
	name          string
	rollno        int
	comp_sc_marks int
	ds_marks      int
}

// Pointer reciever
func (c *Student) new_comp_marks(new_marks int) {
	// NOTE
	// if the struct is very large, better to use Pointer reciever as Value rec. will
	// be making a copy of this. NOT OPTIMIZED
	c.comp_sc_marks = new_marks
}

// value reciever
func (s Student) total_marks() int {
	//it makes a copy of s, and then to all the calculations/changes on it,
	// thus ensuring originnal one stay intact.
	return s.comp_sc_marks + s.ds_marks
}

//normal func
func total_marks_2(s Student, new_marks int) Student {
	s.comp_sc_marks = new_marks
	return s
}

func main() {
	student1 := Student{
		name:          "Sumedha",
		rollno:        151020,
		comp_sc_marks: 99,
		ds_marks:      100}

	student1.new_comp_marks(120)
	fmt.Println(student1.comp_sc_marks)

	// BUT you have to overwrite a student1
	student1 = total_marks_2(student1, 500)
	fmt.Println(student1.comp_sc_marks)

}

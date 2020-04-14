package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// `json`:"name" will put "name" field when creating a json obj,
	// instead of Name
	Name  string `json:"name"`
	Age   int
	Phone int
}

// NOTE1
// The functions in the encoding/json standard package will only encode
// and decode the exported struct fields, which is
// why the first letters of the field names in the above example are all upper cased.

// NOTE2
// 1. When the first letter is capitalised, the identifier is public to any piece of code
// that you want to use.
// 2. When the first letter is lower case, the identifier is private and could only be accesed
// within the package it was declared.

// Exported identifiers - An identifier may be exported to permit
// access to it from another package.

//func main() {
	// json marshalling
	// generating json from a go object

	person_object := Person{Name: "Rahul Verma", Age: 20, Phone: 7018471137}
	person_byte_array, _ := json.Marshal(person_object)

	fmt.Println(string(person_byte_array))
	//{"name":"Rahul Verma","Age":20,"Phone":7018471137}

	// json unmarshalling - converting a byte array to Go object
	var p2 Person

	_ = json.Unmarshal(person_byte_array, &p2)
	fmt.Println(p2.Name)

}

// INDEPTH EXPLANATION
// In Go, package-level visibility is dictated by the case of the name. "Foo"
// is exported, "foo" is not.

// So in this case, encoding/json cannot even see your struct fields. As far
// as it is concerned, they don't exist. As unexported fields, you can only
// access them from within the package in which they are declared (main, in
// this case).

// The solution is to capitalize the names, and use a "struct tag" to provide
// the name to use when encoding/decoding JSON data.

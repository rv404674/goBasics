package main

import "fmt"

// Contains the basics of hashmap and its form in GoLang
// PROS of hashMap - 1. COnstant search time. 2. Arbitary keys.

func main() {
	// MAP

	var idMap map[string]int     // key, val
	idMap = make(map[string]int) //creates a hash map

	//second way
	idMap2 := map[string]int{
		"joe":   12,
		"rahul": 24,
	}

	// accessing the element
	x := idMap2["joe"] //returns zero is joe is not present
	fmt.Println(x)

	//adding a key val pair or change the existing one
	idMap["jame"] = 945

	delete(idMap, "jame")

	// to check is some key exists
	val, p := idMap["jame"]
	// val is value, p is presence
	fmt.Println(val, p)

	for key, val := range idMap {
		fmt.Println(key, val)
	}

}

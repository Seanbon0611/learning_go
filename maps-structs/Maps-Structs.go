package main

import (
	"fmt"
)

func main() {
	//MAPS
	/*
		maps are very similar to objects in that they create key-value pairs
		but the syntax is a bit different than other languages
	*/
	superman := map[string]string{
		"firstName": "Clark",
		"lastName":  "Kent",
		"planet":    "Krypton",
	}
	//to create this object we start with 'map' to initate that it is going to be an object, followed by square brackets.
	//within the square brackets is the data type of the keys, and outside of the brackets is the data type of the values
	fmt.Println(superman)

	//NOTE*** You cannot have a slice as a key in a map BUT you can have an array as a key

	//An alternative way of making a struct is with the make syntax

	batman := make(map[string]string)
	batman = map[string]string{
		"firstName": "Bruce",
		"lastName":  "Wayne",
		"planet":    "Earth",
	}

	//To get values from a specific key from a map by placing the key within square brackets
	fmt.Println(batman["firstName"])

	//adding a new key-value pair to a map
	superman["height"] = "6ft 3in"
	fmt.Println(superman["height"]) //notice how the order is different, the return order of a map is not guarenteed

	//delete entery from a map
	delete(superman, "height")
	fmt.Println(superman)
	//note: if you were to check for superman["height"] again, it would actuallly give you 0 instead of an error
	//to resolve this issue, we use the ", ok" syntax
	_, ok := superman["height"] //"_" is the write-only syntax
	fmt.Println(ok)             // will return false if the key was not found, and will return true if

	//maps are passed by reference so changing one variable that points to a map is going to change the other one.
}

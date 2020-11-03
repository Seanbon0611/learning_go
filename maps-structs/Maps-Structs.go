package main

import (
	"fmt"
	"reflect"
)

//Embedding Examples refer to line 118
type User struct {
	Username string `required max:"10"` //sets that a user must have a username and can be a max of 10 characters
	isAdmin  bool
}

type Post struct {
	User
	Title   string
	Content string
}

//COLLECTION TYPES
type Superhero struct { //refer to line 60 for more info
	firstName string
	lastName  string
	planet    string
	age       int
	powers    []string
}

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

	//STRUCTS
	/*
		Structs are comparable to classes that you would see in languages like JavaScript or Ruby

		If you check out line 9 we are declaring the attributes we're going to include in a struct along with the datatype associated with that attribute.
		in structs we can use arrays and slices as elements tied to that attribute. reference the powers attribute.

		in the below example we are creating a new instance of a superhero named wonderwoman along with the the attributes of that superhero
	*/
	wonderWoman := Superhero{
		firstName: "Diana",
		lastName:  "Prince",
		planet:    "Earth-Two",
		age:       800,
		powers: []string{
			"Super Strength",
			"Telepathy",
			"Astral Projection",
		},
	}
	fmt.Println(wonderWoman)
	fmt.Println(wonderWoman.powers[1]) // to get just one attribute we want from a an instance of a struct we would use the "." followed by the requested attribute

	//There is also something called positional syntax where you actually don't need to list name of the attribute, just the value so long as it follows the struct order
	//This method is not advised as it may be harder to read and you may run into maintainance issues in case new attributes are added into the struct
	theFlash := Superhero{
		"Henry",
		"Allen",
		"Earth",
		300,
		[]string{
			"Super Speed",
			"Super Reflexes",
			"Super Durablilty",
			"Time-travel",
		},
	}
	fmt.Println(theFlash)

	//Anonymous Struct
	/*
		Has the same functionality as a regular struct, but is condensed without a declaration at the package level.
		cannot be used anywhere else however because it is ananoymous

		unlike maps, structs are value types so you can assign a variable to an existing struct instance and it will not overwrite the original instance, it will create a copy of it
		just like arrays, we can use the "&" operator and make it a pointer
	*/
	broTheChihuahua := struct {
		breed string
		name  string
	}{breed: "Chihuahua", name: "Bro"}
	dogeTheShiba := broTheChihuahua
	dogeTheShiba.breed = "Shiba"
	fmt.Println(broTheChihuahua)
	fmt.Println(dogeTheShiba)

	//Embedding
	/*
		As go is a functional language and cannot use inheritance like OOP languages do
		Go uses something like inheritance called composition through embedding.
	*/
	newUser := Post{
		User:    User{Username: "AspiringProgrammer22", isAdmin: false},
		Title:   "My First Post",
		Content: "This is my first post!",
	}
	/*
		alternate way of writing:
		newUser := Post{}
		newUser.Username = "AspiringProgrammer22"
		newUser.isAdmin = false
		newUser.Title = "My First Post"
		newUser.Content = "This is my first post!"
	*/
	fmt.Println(newUser)
	//In the example above it is not showing that a post belongs to a user, they are two seperate structs. But we are embedding the user attributes to a post struct
	//In order to use data interchangably like inheritance, you would need to use something called interfaces which will be went over in a future example

	//TAGS
	//Tags can be used for validations, refer to User struct for an example

	t := reflect.TypeOf(User{})           //get the type of an object that we're working with
	field, _ := t.FieldByName("Username") //grab a field from that type
	fmt.Println(field.Tag)                //Getting the Tag property of the field

	/*
		Again we are using the reflect package and using the TypeOf function passing in the struct,
		we are then assigning the field getting the Username attribute from the USer struct
		then we are printing the tag that is associated with that field we made

		On it's own the tags don't do anything, it would need a validation framework to really do anything
	*/
}

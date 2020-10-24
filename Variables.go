package main

import (
	"fmt"
	"strconv"
)

//declaring variables at the package level(outside the main func) you cannot use the := syntax
var codingLanguage string = "GO"

//At the package level you could also declare a block of variables
//used to organize and have code a little cleaner
var (
	sport        string = "Football"
	position     string = "Quarterback"
	playerName   string = "Jimmy Garappolo"
	playerNumber int    = 10
)

func main() {
	//There are 3 different ways in which you can assign variables
	//1
	//assigning the variable first
	var firstName string
	firstName = "Sean"
	//2
	//assigning the variable and type all on one line
	var lastName string = "Dever"
	var j float32 = 23
	//3
	//implicit declaration
	//letting go determine what datatype it is based on the value you're assigning
	i := 69
	//using the above syntax if you wanted to let GO know you want it to be a float32 instead of it being inferred as as an int, you would add a "." at the end
	//ex: i := 72.
	fmt.Println(firstName, lastName, i)
	//%v prints the value, %T prints the type
	fmt.Printf("%v, %T", j, j)

	//SHADOWING
	//shadowing is when you have a variable with in the package level, as well as that same variable within that function's scope
	//The variable within the inner-most scope will take precidence
	var codingLanguage string = "golang"
	fmt.Println(codingLanguage)
	//even though codingLanguage was declared at the package level, it would print "golang" instead of "GO"

	//CONVERTING DATATYPES
	var a int = 11
	fmt.Printf("%v, %T\n", a, a)

	//we're assigning a the variable b to take a float32 datatype
	//then we're assigning the value of be to a wrapped around a float32 conversion function
	var b string
	b = strconv.Itoa(a)
	fmt.Printf("%v, %T\n", b, b)
	//If you tried to use this same logic converting an int to a string you would need the "strconv"(string conversion) package.
	//Otherwise you would get a character tied to that unicode character
}

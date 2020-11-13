package main

import (
	"fmt"
)

func main() {
	var a int = 42
	/*
		'*' is declaring that it will be a pointer to an integer
		'&' is the "address of" operator that will allow us point to a variable itself instead of just making a copy of it
	*/
	var b *int = &a //assigning variable b to point to the value of variable a
	fmt.Println(a)
	fmt.Println(b) //notice that when run does not hold the value 42 like 'a' does. This is because it is pointing to a's location in memory, not it's value
	//the way we go about this to get the value instead of the address, we need to de-reference where we apply the '*' in front of the variable to go into the address and retrieve the value at that address
	fmt.Println(*b)
	a = 27
	fmt.Println(a, *b) // as you can see, now when we change 'a' it affects our 'b' variable
	//now any changes we make to b will also affect a as b is pointing to a and not just copying it's value

	//Pointer as a datatype
	var nums *myStruct
	nums = &myStruct{number: 23}
	fmt.Println(nums) // when you print this see how there is a '&' next to the value, that lets you know that it is holding that address of the object that has a field with the value in it

	//"new" function
	/*
		with the new function, we are able to do the same to initialize a variable to the pointer of an object
	*/
	var names *myStruct
	names = new(myStruct)  //with the syntax you can initialize with values
	(*names).name = "Sean" //You need the att parenthases inorder to properly use the defreferencing operator as it has more presidence over the '.' operator
	fmt.Println((*names).name)
	/*
		you don't need all this, syntactic sugar version:
		var names *myStruct
		names = new(myStruct)
		names.name = "Sean"
		fmt.Println(names.name)

		works the same, it's implicit that we want the underlying object, not to what it's pointing to.
	*/
}

type myStruct struct {
	number int
	name   string
}

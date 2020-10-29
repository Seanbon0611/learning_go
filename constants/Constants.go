package main

import (
	"fmt"
)

//Enumerated constants
const (
	//go infers the pattern of assignments which is why 'd' and 'e' are able to have a value without an error
	_ = iota //_ is a write only variable that tells the compiler that you're going to generate a value, but you don't care what it is, go ahead and throw it away.
	c
	d
	e
)

//another usecase with enumerable constants is we can assign roles using bytes
const (
	isAdmin          = 1 << iota // << bit-shift to the left that with each following will increment
	isHeadquartes                // bit-shift in 1 place 2^2
	canSeeFinancials             // bit-shift in 1 place 2^4

	canSeeCalifornia // bit-shift in 3 place 2^8
	canSeeFlorida    // bit-shift in 4 place 2^16
	canSeeNewYork
	canSeeBoston
	canSeeWashinton
)

/*
iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
*/

func main() {
	//CONSTANTS

	/*
		Go also has constant variables. These variables are immutable meaning that once they are assigned, cannot be redeclared with a new value.
		ex:
		const randomFact string = "Pteronophobia is the fear of feathers or being tickled by feathers"
		randomFact = "Bananas are curved because they grow closer to the sun"
		This would return an error because we've already declared that the random fact will be and it cannot be overwritten
	*/
	//Constants may also be implicit
	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T\n", a+b, a+b)

	fmt.Printf("%v, %T\n", c, c) //refer to line 7
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	var roles byte = isAdmin | canSeeFinancials | canSeeNewYork //assigning roles to be either isAdmin or canSeeFinancials or canSeeCalifornia
	fmt.Printf("%b\n", roles)
	//now we can check to see if a user is an admin
	fmt.Printf("is Admin? %v\n", isAdmin&roles == isAdmin)                             //bit-mask: if user is an admin, we're going to have the value 1 of that bit set and we're comparing to the isAdmin constant
	fmt.Printf("Can see California? %v\n", canSeeCalifornia&roles == canSeeCalifornia) // becasue canSeeCalifornia is not within the roles we get false
}

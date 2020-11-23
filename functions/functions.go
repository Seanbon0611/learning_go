package main

import (
	"fmt"
)

func main() { //functions start with the func keyword, with every application you make, you must have a main func that has no parameters
	//passing by value, we are assigning a variable that is going to hold a value and we pass that variable as an argument in the function
	var user1 string = "sean"
	greet(user1)
	//Not only are we able to pass by values, but we are able to pass pointers along as arguments in functions
	var user2 string = "foo"
	greetUser(&user2)  // see line 23 how we changing the name param which is pointing to our user2 variable
	fmt.Println(user2) //as it's pointing to the actualy variable and not a copy, we are changing the name of user2 within the greetUSer func

	returnNums := returnSum(1, 2, 3, 4, 5, 6, 7)
	returnSumWithPointer := sumWithPointer(1, 2, 3, 4, 5, 6, 7)
	sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("the sum is", returnNums)
	fmt.Println("sum using pointer is", *returnSumWithPointer)

	divideNums, err := divide(5.0, 3.0)
	if err != nil { //if there is an error it will print the error
		fmt.Println(err)
		return
	}
	//error handling for the above function
	fmt.Println(divideNums) //if there are no errors then it will continue and print out the result

	//Anonymous function
	//because the function has no name, we can't put it within the package level because the function cannot be called without a name
	func() {
		fmt.Println("This is an anonymous function")
	}() //by having the "()" at the end of the function that means we are calling it immediately

	//a good use case for anonymous functions is if you want to invoke a function within another function
	incrNums := intSeq()    //setting a variable that will run and keep track of the counter in memory
	fmt.Println(incrNums()) //first call: 1
	fmt.Println(incrNums()) //second call: 2
	fmt.Println(incrNums()) //third call: 3
	newSet := intSeq()
	fmt.Println(newSet()) //newSet is a new variable so even though we are using the same intSeq func it wont continue off of incrNums and will start off as 0

	//METHODS
	//A method is a function that is executing in a known context, in this example we're using structs but you can use any type of context
	greetUser := greeter{
		greeting: "Hello",
		name:     "Sean",
	}
	greetUser.greet()
	greetUser.pointerGreet()
	fmt.Println("the new name is:", greetUser.name)
}

func greet(name string) { //with paramaters you need to declare the datatype next to the name of the param
	fmt.Printf("Hello, %v\n", name)
}

func greetUser(name *string) {
	fmt.Printf("Hello, %v\n", *name)
	*name = "bar"
}

//just like in other languages, functions can take multiple parameters
func message(msg string, idx int) {
	fmt.Println("The value of the index is", idx)
}

//Variatic Paramteres
//We are able to spread multiple values into a function so long as they are of the same data type
//on line 16 we are able to pass in as many arguments as we want and Go will compile it as a slice
//check line 16 for an example
func sum(values ...int) {
	accumulator := 0
	for _, v := range values {
		accumulator += v
	}
	fmt.Println(accumulator)
}

//with return statement
//instead of just printing the value like the above example, we can return theose values that can be used somewhere else.
func returnSum(values ...int) int {
	accumulator := 0
	for _, v := range values {
		accumulator += v
	}
	return accumulator
}

//We can also use pointers
func sumWithPointer(values ...int) *int {
	accumulator := 0
	for _, v := range values {
		accumulator += v
	}
	return &accumulator
}

//handling errors within functions
func divide(a, b float64) (float64, error) { // we are stating that we are having a 2nd return value that will be of type error
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil //we have to include nil because we stated we are going to return a 2nd value and if there is no errors then we just return nil
}

func intSeq() func() int { //returning a function as a type and that function will return an int
	i := 0
	return func() int {
		i++
		return i
	}
}

//METHODS
type greeter struct {
	greeting string
	name     string
}

//This method is going to get a copy of the greeter object which we are giving the name g in the context of this method
//This method is a value type in such is passed by value where it is creating a literal copy
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

//We can also use a pointer reciever so that we are able to manipulate the underlying data we put in
func (g *greeter) pointerGreet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

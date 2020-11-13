package main

import (
	"fmt"
)

func main() {
	//DEFERRED FUNCTIONS
	/*
		the defer keyword will execute any function thats passed in it after the function has finished it's final statement, but before the function actually returns
		so in the example below it's going to go down line by line, run it. if it sees a defer it will recognize it
		then once it see's it is at the bottom on the main func it will then look back to where the defer keyword is and run it

		if you have multiple defers the way that it actually runs is LIFO(Last In First Out) just like a stack would
	*/
	fmt.Println("start")
	defer fmt.Println("mid") //you'll see this we be the last thing to print
	fmt.Println("end")

	//defer will also hold the value of whatever it is at the time when it is called
	a := "start"
	fmt.Println(a)
	a = "end"
	//the code above will print "start"

	//PANICKING
	/*
		Panic is when our application is unable to continue a function and so it panics because it doesn't know what to do
	*/
	// b, c := 1, 0
	// ans := b / c // 1 cannot be divisable by 0
	// fmt.Println(ans) //when ran will generate a panic for us

	//There are going to be times when go will not be able to catch it and will be in a state where it cannot continue to execute

	//Recover
	/*
		Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")

}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*
	The function g takes the int i, and panics if i is greater than 3, or else it calls itself with the argument i+1. The function main defers a function that calls recover and prints the recovered value (if it is non-nil).
*/

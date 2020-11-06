package main

import (
	"fmt"
)

func main() {
	myArray := [3]string{"Hello", "world", "!"}
	//This is a basic for loop, we are iterating over my array and we are going to print each element in the array
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	//say we wanted to have another variable(in a lot of other languages we typically use j to signify we want to do another loop) we can do so like this
	for i, j := 0, 0; i < len(myArray); i, j = i+1, j+1 { //cannot use i++ or j++ because those are statements, not expressions
		fmt.Println(myArray[i], myArray[j])
	}

	//different example of for loop
	var a int = 0
	for ; a < 3; a++ { //we declared a on the line above so we don't need to do so inline, but we do need the semi-colon
		fmt.Println(a)
	}
	//The biggest difference with this method is the scope, when we do it inline it is within it's own scope, if we do it like above it is in the main func's scope

	//In other languages they have while loops, using a while keyword
	//But go does while loops differently
	b := 0
	for b < 10 {
		fmt.Println("while loops")
		b++
		if b == 2 {
			continue // if we want to continue the loops, as normal functionality is it will stop once condition has been met
		}
		if b == 3 { //breaking a loop using conditionals;
			break
		}

		//Nested for loop
		for i := 0; i <= len(myArray); i++ {
			for j := i + 1; j <= len(myArray); j++ {
				fmt.Println(i * j)
			}
		}

		//Label
		/*
			if you wanted to break out a nested for loop, you wouldn't be able to just use a break statement
		*/
	Loop:
		for i := 0; i <= len(myArray); i++ {
			for j := i + 1; j <= len(myArray); j++ {
				fmt.Println(i * j)
				if i*j >= 3 {
					break Loop // just a regular break will only exit out of the first nearest loop, so by using the label at the top of the loop and adding it next to the break, we are able to break at the top level once the condition has been met
				}
			}
		}
		//Collections with for loops
		var mySlice = []int{1, 2, 3}
		for k, v := range mySlice {
			fmt.Println(k, v) // prints out the index(k) as well as the values(v)
		}

		//If you want just the keys, you would just exclude the value portion
		for k := range mySlice {
			fmt.Println(k)
		}
		//BUT if you only wanted the values, we need to use the write only variable
		//if you just left out the k it would still only give you the key
		for _, v := range mySlice {
			fmt.Println(v) // prints out the index(k) as well as the values(v)
		}
	}
}

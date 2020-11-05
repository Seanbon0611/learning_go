package main

import (
	"fmt"
	"time"
)

func main() {
	//typical structure for an if statement:
	answer := false
	if answer == true {
		fmt.Println("the test is true!")
	} else {
		fmt.Println("the test is false!")
	}
	//Initializer syntax
	inventory := map[string]int{
		"Shirt": 100,
		"Pants": 44,
		"Shoes": 13,
	}
	//If statement that will only print if the conditions are true
	if inv, ok := inventory["Jean"]; ok { //first part of if statement is initializer, the "ok" sytnax is the condition that will give true or false
		fmt.Println(inv)
	}
	//initializer syntax cannot be used outside of it's own scope

	//elseif
	myNum := 30
	target := 40
	if myNum < target {
		fmt.Println("Number is too small!")
	} else if myNum > target {
		fmt.Println("Number is too large")
	} else {
		fmt.Println("That is the number!")
	}

	//Switch statements
	/*
		Just like in a lot of other languages, switch statements are used if we may run into a lot of different possible paths
		this would be a lot cleaner than if we were to do a lot of else if statements but has the same functionality.
	*/
	date := time.Now().Weekday() //getting weekday from the "time" package

	//Dpending on which day of the week it is, it will check the cases and returns the one that matches.
	//the default at the bottom will be the option if none of the cases match
	switch date {
	case time.Monday:
		fmt.Println("It is Monday")
	case time.Tuesday:
		fmt.Println("It is Tuesday")
	case time.Wednesday:
		fmt.Println("It is Wednesday")
	case time.Thursday:
		fmt.Println("It is Thursday")
	case time.Friday:
		fmt.Println("It is Friday")
	case time.Saturday:
		fmt.Println("It is Saturday")
	case time.Sunday:
		fmt.Println("It is Sunday")
	default:
		fmt.Println("I've lost track of the days! x_x")
	}

	//Tagless switch statement
	//He we're setting a variable, the switch statement is looking at the cases as there are no tags and determining which to return based on the conditionals in the cases
	//In GO the break keyword is implied and it will return the first one once a condition has been met unline other languages
	i := 10

	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
	case i <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than 20")
	}
}

package main

import "fmt"

func main() {
	/*
		in front of the function invocation we place go to make it a goroutine
		This will tell go to spin off a "green thread" and run this funtion in the green thread.
		Most other programming languages use Operating system threads, this means they have an individual call stack dedicated to the execution of the code within that thread.
		In Go green threads are abstractions of a thread called goroutines.
		Inside of the go runtime there is a scheduler that is going to map the goroutines onto OS threads for periods of time.
		and the scheduler takes turns with every CPU thread that is avaiible and allocate a certain amount of processing time for these threads. But we dont have to handle these low level threads, just the larger goroutines.
	*/

	go sayHello()
}

func sayHello() {
	fmt.Println("Hello!")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		in front of the function invocation we place go to make it a goroutine
		This will tell go to spin off a "green thread" and run this funtion in the green thread.
		Most other programming languages use Operating system threads, this means they have an individual call stack dedicated to the execution of the code within that thread.
		In Go green threads are abstractions of a thread called goroutines.
		Inside of the go runtime there is a scheduler that is going to map the goroutines onto OS threads for periods of time.
		and the scheduler takes turns with every CPU thread that is avaiible and allocate a certain amount of processing time for these threads. But we dont have to handle these low level threads, just the larger goroutines.
	*/
	//Tells the main function to create another goroutine
	// go sayHello()

	//example with closures
	var msg string = "Greetings, Planet!"
	go func() { //anonymous function that is printing out the msg variable we declared above.
		fmt.Println(msg) //Even though we are within the scope of the anon function, we have access to the msg variable thats within the main func's scope
	}()
	msg = "Goodbye"                    // when we run the file we'll see it will print out "Goodbye". This is because the main function is executing before the goroutine is able to execute
	time.Sleep(100 * time.Millisecond) //if we don't have this, we will never see the sayHello func get executed because it doesn't have enough time to because that is the end of the main func

}

func sayHello() {
	fmt.Println("Hello!")
}

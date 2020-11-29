package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Having the time.Sleep is not the best practice though so we will use the sync package to create a WaitGroup
//What a Waitgroup does is syncronize multiple goroutines together
//We are using this variable to syncronize the main function with the goroutine
var wg = sync.WaitGroup{}
var counter int = 0

//Mutex
//A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer.
//any amount of goroutines can read the data, but only one can write at a time, if anything is reading then nothing can be wrote at all
var m = sync.RWMutex{}

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

	wg.Add(1) //Adding the following goroutine below
	//example with closures
	var msg string = "Greetings, Planet!"
	go func() { //anonymous function that is printing out the msg variable we declared above.
		fmt.Println(msg) //Even though we are within the scope of the anon function, we have access to the msg variable thats within the main func's scope
		wg.Done()        //method that tells the Waitgroup that the goroutine is done, by decrement the number of waitgroups
	}()
	msg = "Goodbye" // when we run the file we'll see it will print out "Goodbye". This is because the main function is executing before the goroutine is able to execute
	//once the goroutine is done we exit the application by using the Wait method
	wg.Wait()

	runtime.GOMAXPROCS(100) //sets the maximum amount of threads
	for i := 0; i < 10; i++ {
		wg.Add(2)
		//we are locking before each goroutine and then unlocking within the goroutine itself
		m.RLock() //RLock locks rw for reading
		go sayHello()
		m.Lock() //Write lock
		go increment()
	}
	wg.Wait()

}

//SYNCHRONIZING MULTIPLE GOROUTINES OUTSIDE OF MAIN FUNC

func sayHello() {
	fmt.Printf("Hello! %v\n", counter)
	m.RUnlock() //unlocks the Rlock
	wg.Done()
}

func increment() {
	counter++
	m.Unlock() //Unlocks the write lock
	wg.Done()
}

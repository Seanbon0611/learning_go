package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Hello")
	//when creating a channel, you have to use the make function in order to allow the runtime to setup the channel
	ch := make(chan int) //chan keyword signfies we are making a channel and the data type that channel will take
	//adding 2 goroutines
	wg.Add(2)
	//first goroutine recives data from the channel
	go func() {
		i := <-ch //recieving data from the channel and assigning it to the variable "i"
		fmt.Println(i)
		wg.Done()
	}()
	//goroutine thats adding data to a channel
	go func() {
		ch <- 42 //arrow "<-" declares that we're putting the data 42 into the channel
		wg.Done()
	}()
	wg.Wait()
}

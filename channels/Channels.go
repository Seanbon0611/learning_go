package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
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

	//Asyncronous Channel

	new_ch := make(chan int)
	//creating 5 goroutines
	for j := 0; j < 5; j++ {
		wg.Add(2)
		//reciever
		go func() {
			i := <-new_ch
			fmt.Println(i)
			wg.Done()
		}()
		//sender
		go func() {
			new_ch <- 69 //goroutine gets paused here UNTIL space is available in the channel
			wg.Done()
		}()
	}
	wg.Wait()

	//improved dataflow with goroutine
	ch_3 := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { //setting a recieve only channel as a parameter for this function
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch_3)
	go func(ch chan<- int) { //sending data into the channel with a channel set as the parameter and that the channel can only take in an int
		ch <- 11
		wg.Done()
	}(ch_3)
	wg.Wait()
}

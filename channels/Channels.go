package main

import (
	"fmt"
	"sync"
	"time"
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

	//Polymorphic Channels
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

	//Buffered Channels
	/*
		What buffers does is it when your sender or reciever needs a little time to process and prevents it from blocking the other side because of the delay
	*/
	ch_4 := make(chan int, 50) //the 2nd parameter here is the buffer, it tells go to create a channel that has an internal datastore that can store 50 integers
	wg.Add(2)
	go func(ch <-chan int) {
		//ranging over channel to pull the single value from the channel each
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch_4)
	go func(ch chan<- int) {
		ch <- 22
		ch <- 18
		close(ch_4) //we are using the close function to prevent a deadlock, closing the channel signifies that once the data is pushed into the channel that there will be no more data to be pushed and we are closing the channel so the range knows when the channel ends
		wg.Done()
	}(ch_4)
	wg.Wait()

	//Select Statements
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:05:05"), entry.severity, entry.message)
	}
}

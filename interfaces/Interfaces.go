package main

import "fmt"

//Interfaces are a type just like structs
//Interfaces describe behaviors, unlike structs which describe data
type Writer interface {
	Write([]byte) (int, error) //writing a slice of bytes that will return an integer and an error
}

type ConsoleWriter struct { //initiating an empty struct
}

//method that gets a copy of the ConsoleWriter struct getting data which is a slice of bytes and we are expected to return an int and error
//we are getting the data within the slice, converting it to type string, then printing and returning.
func (cr ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Println(data)
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	var w Writer = ConsoleWriter{}   //assigning variable "w" to be of Type Writer(Our Interface) and to equal our ConsoleWriter struct
	w.Write([]byte("Hello, World!")) //We are calling the Write function which we've made above and is also listed on our interface passing in Hellow World converted into a slice of bytes

	myInt := Counter(0)          //Creating an instance of the Counter starting at 0
	var inc Incrementer = &myInt //creating a variable with the type incrementer that will point to the myInt Counter
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment()) //as long as i < 10 run the inrement function that will add to the counter
	}
}

type Incrementer interface {
	Increment() int
}

type Counter int

func (c *Counter) Increment() int {
	*c++
	return int(*c)
}

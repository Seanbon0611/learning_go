package main

import (
	"bytes"
	"fmt"
)

//Interfaces are a type just like structs
//Interfaces describe behaviors, unlike structs which describe data
type Writer interface {
	Write([]byte) (int, error) //writing a slice of bytes that will return an integer and an error
}

//interface that has a close method that returns an error in case something happens
type Closer interface {
	Close() error
}

//Interface thats composed of other interfaces
//
type WriterCloser interface {
	Writer
	Closer
}

type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

type ConsoleWriter struct { //initiating an empty struct
}

//Takes in some data, prints it out to the console in 8 character increments
func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) //gets data and stores into the buffer
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 { // as long as the buffer has more than 8 characters
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
	}
	return n, nil
}

//Gets whats leftoveer from the buffer(remainding characters that don't meet the 8 char requirement)
//and prints that to the console until the buffer is empty
func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

//constructor function that returns a pointer to a BufferedWriterCloser, This initializes internal buffer to a new buffer
func NewBufferedWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

//method that gets a copy of the ConsoleWriter struct getting data which is a slice of bytes and we are expected to return an int and error
//we are getting the data within the slice, converting it to type string, then printing and returning.
// func (cr ConsoleWriter) Write(data []byte) (int, error) {
// 	fmt.Println(data)
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

func main() {
	// var w Writer = ConsoleWriter{}   //assigning variable "w" to be of Type Writer(Our Interface) and to equal our ConsoleWriter struct
	// w.Write([]byte("Hello, World!")) //We are calling the Write function which we've made above and is also listed on our interface passing in Hellow World converted into a slice of bytes

	myInt := Counter(0)          //Creating an instance of the Counter starting at 0
	var inc Incrementer = &myInt //creating a variable with the type incrementer that will point to the myInt Counter
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment()) //as long as i < 10 run the inrement function that will add to the counter
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Greetings Everyone, My Name is Sean and i'm a web developer!")) //converts string to a byte slice
	wc.Close()

	//type conversion
	bwc := wc.(*BufferWriterCloser) //we get the object then use the "." operator
	fmt.Println(bwc)                //instead of the writer closer, becomes a bufferwritercloser

	r, ok := wc.(*BufferWriterCloser) //we get the object then use the "." operator
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion Failed")
	}

	//Empty Interface
	//An in terface that is defined on the fly without any methods on it
	//Everything can be cast into an object that has no methods on it, even Primitives
	var myObj interface{} = NewBufferedWriterCloser()
	fmt.Println(myObj)

	//Interface Type Switch
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an Integer")
	case string:
		fmt.Println("i is a String")
	default:
		fmt.Println("i is an unknown type")
	}
	/*
		Rules & Best Practices with Interfaces
		-----------------------------------------
		- Use many small interfaces
			it's better to have many smaller interfaces than one big one becas\use single method interfaces are more powerful and flexible
				ex: io.Writer, io.Reader, interface{}
		- Dont export interfaces for types that will be consumed
			if you don't need to export the interface yourself, then don't
		- Do export interfaces for types that will be used by package
			if you're going to be pulling a value in, accept an interface instead of a concrete type because go has implicit implementation
		- Design functions and methods to reciveve interfaces whenever possible
	*/
}

type Incrementer interface {
	Increment() int
}

type Counter int

func (c *Counter) Increment() int {
	*c++
	return int(*c)
}

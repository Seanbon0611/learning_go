package main

import (
	"fmt"
)

func main() {
	//BOOLEANS
	//Everytime you initialize a varaible in GO, it has a "0" value
	var a bool
	fmt.Printf("%v, %T\n", a, a)
	//this print function shows that the value is == false because the a variable is being stored with a value of 0, and 0 is falsey

	//NUMERIC DATATYPES
	//int is an unspecified size, being at least a 32bit, but could be 64 or 128 depending on your system
	b := 42
	fmt.Printf("%v, %T\n", b, b)
	/*
		there are 8bit integers "int8" which can range from -128 - 127
		16bit integers "int16" which can range from -32,768 - 32,767
		32bit integers "int32" ranges from -2,147,483,648 - 2,147,483,647
		64 bit integers "int64" ranges from -9,223,372,036,854,775,808 - 9,223,372,036,854,775,807

		Unsigned integers can hold a larger positive value, and no negative value.
		uint8 can range from 0 - 255
		unint16 range from 0 - 65,535
		uint32 range from 0 - 4,294,967,295

		When you are doing mathematical equations, they must be of the same datatype
		ex:
		var a int = 10
		var b int8 = 3
		you cannot do a + b because they are two seperate datatypes, but you can do a type conversion to make it work
	*/

	// BITWISE OPERATOR
	// there are 4 bit operators: "&", "|", "^", and "&^"
	var c int = 10      //value in binary 1010
	var d int = 3       //value in binary 0011
	fmt.Println(c & d)  //Binary AND Operator copies a bit to the result if it exists in both operands.
	fmt.Println(c | d)  //Binary OR Operator copies a bit if it exists in either operand.
	fmt.Println(c ^ d)  //Binary XOR Operator copies the bit if it is set in one operand but not both
	fmt.Println(c &^ d) //Binary AND NOT copies a bit if it exists doesn't in either operand
	//Bit shifting
	e := 8
	fmt.Println(e << 3) //shift the bit left 3 plades
	fmt.Println(e >> 3) ////shift the bit left 3 places

	//FLOAT DATATYPES
	//Floats can be either float32 or float64
	var f float32
	f = 3.14
	fmt.Printf("%v, %T\n", f, f)

	g := 13.7e72
	fmt.Printf("%v, %T\n", g, g)

	//COMPLEX NUMBER DATATYPE
	//Complex can either be complex64 or complex 128
	//the i in 2i stands for imaginary(?)

	var h complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(h), real(h))
	fmt.Printf("%v, %T\n", imag(h), imag(h))

	//STRINGS
	/*Strings are in the UTF-8 format, or are recognized as bytes
		They could also be UTF-32
	Ex: "s" in UTF-8 is 115
	*/
	s := "s"
	o := []byte(s) //byte-slice
	fmt.Printf("%v, %T\n", o, o)

}

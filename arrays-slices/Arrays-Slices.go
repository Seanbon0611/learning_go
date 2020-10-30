package main

import (
	"fmt"
)

func main() {
	/*
		When declaring an array, we start with the square brackets with the size of the array within those brackets followed by the datatype that array can store.
		If you want to initialize the array with values next to the datatypes you place curly brackets with the values.

		When initiating arrays with values, you don't need to have in index of the array within the square brackets. instead you can fill it with [...] which means set the array length equal to the amount of elements you're initializing with
	*/
	grades := [...]int{92, 85, 13}
	fmt.Printf("%v\n", grades)

	var students [3]string
	students[0] = "Sean"
	students[1] = "Elaina"
	students[2] = "Jon"
	fmt.Printf("%v\n", students)
	//displaying the SIZE of the array not the amount of students within the array
	//len function returns the length/size
	fmt.Printf("Number of Students: %v\n", len(students))

	//Array of arrays
	var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	/*
		could also be written as:
		var matrix [3][3]int
		matrix[0] = [3]int{1,0,0}
		matrix[1] = [3]int{0,1,0}
		matrix[2] = [3]int{0,0,1}
	*/
	fmt.Println(matrix)
	//In Go, arrays are considered values
	a := [...]int{1, 2, 3}
	//When you copy an array, you area creating a literal copy not overwriting.
	b := a
	b[1] = 44
	fmt.Println(a)
	fmt.Println(b)
	//If you want to overwrite the original array, you would need to use pointers
	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)
	//d points to c, so when we write 2nd element of d to be 5, since d points to c it is changing c's 2nd element

	//Slice
	// with a slice we are initiating as a literal with square brackets
	e := []int{1, 2, 3}
	f := e
	f[1] = 12
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e)) //num of elements in the slice does not match the size of the backing array because the slice is the projection of the underlying array
	//With slices you do not need the pointer, as you can see when we are assigning f to e and change f[1] to 5, it is also changing e[1] to 5
	//Other methods to create slices
	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := g[:]   //slice of all elements
	i := g[3:]  //slice from the 4th element to the end
	j := g[:6]  //slice the first 6 elements
	k := g[3:6] //slice elements 4-6
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	//MAKE FUNCTION
	/*
		Takes 2 or 3 arguments
		the first arugment is the tpye of object you want to create, the 2nd argument is the length of the slice, the 3rd argument is the capacity
		why? Because unlike arrays, slices don't have to have a fixed size, we can add or remove elements from them.
	*/
	//When creating a slice, everything is getting set to a '0' value.
	l := make([]int, 3, 100)
	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("Capactiy: %v\n", cap(l))
	//if you want to add to the slice, we would need to use the append function
	m := []int{}
	m = append(m, 1)
	//variatic function: everything after the first argument is going to be interpreted as a value to append to the slice passed into the first argument
	m = append(m, 2, 3, 4, 5)
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capactiy: %v\n", cap(m))

	//stack operations
	n := []int{1, 2, 3, 4, 5}
	o := append(n, 6) //adds to the end of a slice
	p := n[1:]        //shifts a slice(remove first element)
	q := n[:len(n)-1] //pops the last element in a slice
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	/*to remove elements in the middle of a slice, we have to concat two slices together
	the first slice goes up to the point of the element we want to remove
	the second slice is going to hold the elemts after the elemnt we want to remove
	we then use the spread operator "..." to spread thin gs out to work with the append function
	*/
	r := append(n[:2], n[3:]...)
	fmt.Println(r)
}

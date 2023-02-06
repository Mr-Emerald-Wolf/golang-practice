package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 25

	var ptr = &myNumber
	fmt.Println("Value of pointer is: ", ptr)
	fmt.Println("Value of object at pointer is: ", *ptr)
	
	*ptr = *ptr * 2 
	fmt.Println("New Value of object at pointer is: ", myNumber)

}

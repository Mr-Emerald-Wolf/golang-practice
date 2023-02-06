package main

import "fmt"

func main()  {
	fmt.Println("Print an array in golang")

	var fruits [3]string
	fruits[0] = "Tomato"
	fruits[1] = "Orange"
	fruits[2] = "Mango"

	fmt.Println("Fruits: ", fruits)
	fmt.Println("No. of fruits: ", len(fruits))

	// Direct initialization of array
	var snacks = [3]string{"chips", "cookies", "icecream"}
	fmt.Println("Snacks: ", snacks)
}
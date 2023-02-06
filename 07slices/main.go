package main

import "fmt"
func main()  {
	fmt.Println("Hello world wassup bros")
	var snacks = []string{"chips", "cookies", "icecream"}
	fmt.Printf("Type of object %T\n", snacks)
	snacks = append(snacks, "Chocolate", "Nachos")
	fmt.Println(snacks)
	snacks = append(snacks[1:3], "Chocolate", "Nachos")
	fmt.Println(snacks)

}
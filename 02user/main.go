package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	rating := "Enter the rating for our Pizza: "
	fmt.Println(rating)
	// comma ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
}

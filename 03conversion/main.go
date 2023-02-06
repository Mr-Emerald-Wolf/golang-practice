package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main()  {
	fmt.Println("Welcome to our Pizza App \nPlease rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Thanks for rating:", numRating + 1)
}
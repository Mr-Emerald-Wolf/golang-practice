package main

import "fmt"

type User struct { 
	Name string 
	Age int 
	Status bool
}

func add(x int, y int) int {
	return x + y
}

func (u User) setAge(x int) { 
	u.Age = x
	fmt.Println(u.Age)
} 

func main() {
	fmt.Println("Making a struct: ")
	rohan := User{"Rohan", 20, true}
	fmt.Printf("User info: %+v \n", rohan )
	fmt.Println(add(4,5))
	rohan.setAge(50) // Only copy of object is passed to the method 
	fmt.Println(rohan.Age)
}

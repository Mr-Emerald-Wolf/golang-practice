package main

import "fmt"
// := declaration not allowed outside main function 

const LoginToken = 10000;

func main() {
	fmt.Println("Hello from GO")
	var username string = "rohan"
	fmt.Printf("hello %s\n", username)
	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn)
	var smalLVal uint8 = 255;
	fmt.Println(smalLVal)

	// implicit style 
	var text = "hello";
	fmt.Println(text)
	// no var style 
	total := 3000.0
	fmt.Printf("Total: %T: %f\n", total, total)
	fmt.Println(LoginToken)

}



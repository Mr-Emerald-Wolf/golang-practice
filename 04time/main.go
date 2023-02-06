package main

import (
	"fmt"
	"time"
)
func main()  {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate  := time.Date(2020, time.April, 10, 23, 0, 40, 50, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))


}
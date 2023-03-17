package main

import "fmt"

func main() {

	//defer statement is used when we want the execution of that code fragment at the end of execution
	defer fmt.Println("world")
	fmt.Println("hello")
}

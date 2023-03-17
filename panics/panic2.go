package main

import "fmt"

func main() {
	//panic and defer is used together when program cannot continue execution
	//and it should be recovered from that state
	defer fmt.Println("here is the defer function")
	panic("divide by 0 error")
}

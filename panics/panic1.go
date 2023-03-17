package main

import "fmt"

func divideByZero() {
	// Use this deferred function to handle errors.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HERE")
			fmt.Println(err)
			fmt.Println(0)
		}
	}()
	// Cause an error.
	a := 0
	b := 10 / a

	//my personal analysis
	//10/0 will not work with panic and it will throw error because go is compile time language
	//so at compile time it will be thrown but our defer function runs at the end of the program and will handle run time error

	fmt.Println(b)
}

func main() {
	// Create a divide by zero error and handle it.
	divideByZero()
}

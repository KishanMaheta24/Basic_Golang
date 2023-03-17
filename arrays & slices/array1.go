package main

import "fmt"

//Var array_name[length]Type - uninitialized array
// var arr = [len] {values...} initialized array
//short hand declaration

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var arr1 = [10]int{1, 2}
	fmt.Println(arr)
	fmt.Println(arr1)
}

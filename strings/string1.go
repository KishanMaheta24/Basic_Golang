package main

import "fmt"

func main() {

	arr := []rune{'a', 'b', 'c', 'd'}
	arr1:="asdhjad"

	for _, obj := range arr {
		// fmt.Print(obj," ")
		fmt.Printf("char: %c\n",obj)
		fmt.Printf("char : %#U",obj)
	}
	fmt.Println(len(arr1))
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	arr := []rune{'a', 'b', 'c', 'd'}
	arr1 := "asdhjad"
	scanner:=bufio.NewReader(os.Stdin)
	val,_:=scanner.ReadString('\n');
	for _, obj := range arr {
		// fmt.Print(obj," ")
		fmt.Printf("char: %c\n", obj)
		fmt.Printf("char : %#U", obj)
		// fmt.Printf("val: %#v",obj)
	
	}
	fmt.Println(len(arr1),val)
}

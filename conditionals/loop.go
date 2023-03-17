package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//will not work
	//sum := 0
	//for i := 0; i < 10; (++i) {
	//	sum += i
	//}
	//fmt.Println(sum)

	//while loop
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	//infinite loop
	//for{
	//
	//}

}

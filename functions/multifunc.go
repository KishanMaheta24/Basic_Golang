package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}
func main() {
	a, b := swap(4, 5)
	fmt.Println(a, b)
}

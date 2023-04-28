package main

import (
	"fmt"
	"math"
)

func main() {
	x := 100

	if x == 50 {
		fmt.Println("ABC")
	} else if x == 100 {
		fmt.Println("DEF")
	} else {
		fmt.Println("GHI")
	}


	if x := math.Sqrt(100); x == 10 {
		fmt.Println("Germany")
	}

}

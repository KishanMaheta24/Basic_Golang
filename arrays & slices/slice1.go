package main

import (
	"fmt"
)

var mat []int

var mat1[][]int

func main() {

	for i := 0; i < 5; i++ {
		mat = append(mat, i)
	}

	fmt.Println(mat)

	for i := 0; i < 3; i++ {

		var temp[]int
        for j := 0; j < 4; j++ {
      temp=append(temp, i+j)
        }
		mat1=append(mat1, temp)
    }
	fmt.Println(mat1)
}

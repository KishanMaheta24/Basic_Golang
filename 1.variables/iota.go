package main

import "fmt"

const (
	first  = iota
	second = iota
	third  = iota
)

// #2

// const (
// 	first = iota + 1
// 	second
// 	third
//  )


// #3

// const (
// 	first = iota + 1
// 	second
// 	_
// 	fourth
//  )

func main() {
	fmt.Println(first, second, third)
}

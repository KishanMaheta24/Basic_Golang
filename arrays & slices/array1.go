package main

import "fmt"

//Var array_name[length]Type - uninitialized array
// var arr = [len] {values...} initialized array
//short hand declaration


//1
// array_name:= [length]Type{item1, item2, item3,...itemN}

//2
// arr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
// {"C++", "Go", "HTML"}}


//3
// x := [...]int{10, 20, 30}

//4
// 	x := [5]int{1: 10, 3: 30}

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var arr1 = [10]int{1, 2}
	fmt.Println(arr)
	fmt.Println(arr1)
}

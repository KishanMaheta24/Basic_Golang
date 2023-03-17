package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Id    int
	Marks Marks
}

type Marks struct {
	Maths     int
	Physics   int
	Chemistry int
}

func main() {
	marks := Marks{90, 80, 70}
	v := Student{"Dummy", 2, 77, marks}
	fmt.Println(v)
	fmt.Println(v.Marks.Maths)
}

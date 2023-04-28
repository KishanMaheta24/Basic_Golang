package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Id    int
	Marks Marks
}

type Marks struct {
	maths     int //private declaration which can only be used by parent struct
	physics   int
	chemistry int
}

type Student1 struct {
	Name string
	Age  int
	Id   int
	Marks
}

func main() {
	marks := Marks{90, 80, 70}
	v := Student{"Dummy", 2, 77, marks}
	v1 := Student1{"Dummy", 2, 77, Marks{90, 80, 70}}
	fmt.Println(v)
	fmt.Println(v.Marks.maths)
	fmt.Println(v1.maths)

}

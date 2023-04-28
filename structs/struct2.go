package main

import "fmt"

type Marks struct {
	Maths     int
	Physics   int
	Chemistry int
}

func (m Marks) getAvg() {

	res:=m.Chemistry+m.Physics+m.Maths;
	fmt.Println(res/3)
}

func main() {
	obj:=Marks{90,80,85}
	obj.getAvg()
}
package main

import "fmt"

type StudentInfo struct {
	Name string
	Age  int
	Id   int
}

func main() {
	v := StudentInfo{"Dummy", 2, 77}
	obj := StudentInfo{Name: "kishan"} //called as struct literals, by using this, it will assign all the not assigned values by default values
	p := &v                            //pointer to structs`
	v.Name = "dummy1"
	fmt.Println(v.Name, v.Age)
	fmt.Println(p.Name)
	fmt.Println(obj)
}

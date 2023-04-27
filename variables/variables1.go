package main

import (
	"fmt"
	"time"
)

//there is no such concept of uninitialized variables in go
//var temp - not allowed
//a:=10 not allowed outside out of main
//var temp int - initialized as 0

var number int
var str string
var num = 1
var str1 = "string here"

var x,y,z int=10,11,12

var(
	name="Go"
	age=5
	isPresent=true
)

func main() {
	//num2:=10 shorthand for declaring
	//good practise to use := if we have value already present
	a:=time.Now()
	fmt.Println("Hello World",a)
	fmt.Print(number," ",str," ",num," ",str1," ",name," ",age," ",isPresent," ",x,y,z)
	
}

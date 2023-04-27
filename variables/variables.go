package main

import "fmt"

/*
Variable names must begin with a letter or an underscore. A variable name can include a to z, A to Z, digits 1 to 9, and '_'.
A variable cannot begin with a number.
A variable name cannot use Go keywords as variable names.
A variable name can only contain alphanumeric characters and underscore.
A variable name cannot contain spaces.
*/

//declaration type 1

var num int = 20           // int variable
var amount float32 = 49.99 // float variable
var isValid bool = true    // bool variable
var name string = "avbcx"  // string variable

//declaration type 2

//var num = 20
//var amount = 49.99
//var isValid = true
//var name = "asdv"

//declaration type 3

//var (
//num = 20
//amount = 49.99
// isValid = true
// name = "asdv"
//)

func main() {

	//declaration type 3 : shorthand declaration

	//num := 20
	//amount := 49.99
	//isValid := true
	//name := "fdg"
	fmt.Println(num, amount, isValid, name)
}

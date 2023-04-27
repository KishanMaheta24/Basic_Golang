package main

import "fmt"

func main() {

	mp := make(map[string]string)
	mp1 := map[string]string{}
	// var map2 map[string]string
	

	mp["fruit"] = "apple"
	mp["bird"] = "peacock"
	fmt.Println(mp["fruit"])
	mp1["dummy"] = "sdad"
	fmt.Print(mp1["dummy"])

}

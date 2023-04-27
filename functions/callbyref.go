package main

import "fmt"

func main() {
   var a int = 100
   var b int = 200

   fmt.Printf("value of a : %d\n", a )
   fmt.Printf("value of b : %d\n", b )

   swap(&a, &b)

   fmt.Printf("After swap, value of a : %d\n", a )
   fmt.Printf("After swap, value of b : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    
   *x = *y    
   *y = temp    
}
package main

import (
	"fmt"
)

func main() {
	x := 10
	changeX(&x)
	fmt.Println(x)

	//nil pointers

	var ptr *int

	fmt.Printf("The value of ptr is : %x\n", ptr)
}

func changeX(x *int) {
	*x = 0
}

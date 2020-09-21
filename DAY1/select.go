package main

import "fmt"

func main() {

	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Print(" the value ", i1, "is 10 \n")
	case c2 <- i2:
		fmt.Print(" the value is 20 \n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Print(" val", i3, "from c3 \n")
		} else {
			fmt.Print(" connection closed")
		}
	default:
		fmt.Print(" no connection ")

	}
}

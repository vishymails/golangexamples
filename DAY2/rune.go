package main

import "fmt"

func main() {
	const val1 = 'Ñ'
	const val2 = 'a'
	fmt.Println("(int32) val1", val1)
	fmt.Println("(int32) val1 %s", val1)
	fmt.Println("(int32) val1", val2)

	odd := [6]int{2, 4, 6, 8, 12, 10}

	var slice []int = odd[1:4]

	fmt.Print(slice)
}

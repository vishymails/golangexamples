package main

import "fmt"

func main() {

	x := 9

	if x == 9 {
		fmt.Println("x value is 9")
	}

	if y := 40; y == 40 {
		fmt.Println("y value is 40")
	} else {
		fmt.Println("x value is differnt")
	}

	z := 70
	if z == 70 {
		fmt.Println("z value is 70")
	} else if z == 90 {
		fmt.Println("z value is 90")
	} else if z == 190 {
		fmt.Println("z value is 190")
	} else {
		fmt.Println("z value is differnt")
	}

	fmt.Print("Enter some number ")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)

	if input%2 == 0 {
		fmt.Printf("is even \n")
	} else {
		fmt.Printf("is odd \n")
	}

}

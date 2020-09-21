package main

import "fmt"

func main() {
	var input int

Loop:
	fmt.Println("You are not eligible for variable pay")
	fmt.Println("Enter your achievements")
	fmt.Scanln(&input)
	if input <= 20 {
		goto Loop
	} else {
		fmt.Print("You are eligible for Variable Pay")
	}
}

package main

import "fmt"

func main() {
	const (
		LosAngeles = 1984 + (iota * 4)
		seoul
		barcelona
		atlanta
		sydney
		athens
		beijing
	)

	fmt.Printf("%-18s %-18v \n", "atlanta", atlanta)

	fmt.Printf("%-18s %-18v \n", "sydney", sydney)

	fmt.Println("", sydney)
}

package main

import "fmt"

func do(steps ...string) {
	defer fmt.Println("All done")

	for _, s := range steps {
		defer fmt.Println(s)
	}
}

func do1(steps ...string) {
	defer fmt.Println("All done")

	defer func() {
		fmt.Println("Hurry u r late")
	}()

	for _, s := range steps {
		fmt.Println(s)
	}
}

func main() {

	do(
		"Accept Resume",
		"find keywords",
		"match the percentage",
		"check percentage",
		"call for an interview",
	)

	do1(
		"Accept Resume",
		"find keywords",
		"match the percentage",
		"check percentage",
		"call for an interview",
	)
}

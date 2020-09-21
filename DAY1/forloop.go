package main

import (
	"fmt"
	"time"
)

func main() {
	for a := 0; a < 11; a++ {
		fmt.Println(a)
	}

	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

	for x := 0; x < 3; x++ {
		for y := 3; y > 0; y-- {

			fmt.Print(x, " ", y, "\n")
		}

	}

	/*
		for true {
			fmt.Print ("infinite")
		}
	*/

	nums := []int{2, 4, 5, 8, 3, 6}
	sum1 := 0

	kvs := map[string]string{"1": "Apple", "2": "Lenovo", "3": "HP"}

	for _, value := range nums {
		sum1 += value
	}

	fmt.Println("acc : ", sum1)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index : ", i)
		}
	}

	for k, v := range kvs {

		fmt.Printf("%s -> %s \n ", k, v)

	}

	for i, c := range "Vishwanath Rao" {
		fmt.Println(i, c)
	}

	for j := 0; j < 10; j++ {

		if j == 0 {
			continue
		}
		fmt.Println("Value : ", j)
	}

	for j := 0; j < 10; j++ {

		if j == 5 {
			continue
		}
		fmt.Println("Value1 : ", j)
	}

	timer := time.NewTimer(time.Second * 9)

	for {
		fmt.Println("Inside infinite loop")

		<-timer.C
		break
	}
}

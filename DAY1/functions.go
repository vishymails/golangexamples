package main

import "fmt"

func fun() int {
	return 900
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0

	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}

	return finalAddValue, finalSubValue
}

type Employee struct {
	fname string
	lname string
}

func (emp Employee) fullName() {
	fmt.Println(emp.fname + emp.lname)
}

func main() {
	e1 := Employee{"Vishwanath", "Rao"}
	e1.fullName()

	x := fun()
	fmt.Println(x)

	fmt.Println(addAll(20, 30, 40, 50, 60, 70))

	fmt.Println(addAll(20, 30, 40, 50, 60, 70, 54, 32, 22))

	fmt.Println(factorial(9))

	//clousure

	number := 10
	squareNum := func() int {
		number += number
		return number
	}

	fmt.Println(squareNum())
	fmt.Println(squareNum())

	fmt.Println(squareNum())

	fmt.Println(squareNum())

}

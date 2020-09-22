package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math : you passed negative number ")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := Sqrt(-32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result1, err1 := Sqrt(32)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

}

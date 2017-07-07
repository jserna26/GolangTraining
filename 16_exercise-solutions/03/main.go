package main

import "fmt"

/*
  Write a function with one variadic parameter that finds the greatest number
  in a list of numbers.
*/
func variadicFunc(x ...int) int {
	var max int
	for _, v := range x {
		if v >= max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(variadicFunc(1, 2, 3, 4, 5))
	fmt.Println(variadicFunc(-10, 2, 70, 69, 70, 101, 55))
}

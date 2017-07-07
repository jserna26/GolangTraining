package main

import "fmt"

/* Modify the previous program (01) to use a func expression. */

func main() {

	//func expression
	myFunc := func(x int) (float64, bool) {
		var even bool

		if x%2 == 0 {
			even = true
		} else {
			even = false
		}
		y := float64(x)

		return (y / 2), even
	}

	fmt.Println(myFunc(2))
	fmt.Println(myFunc(4))
	fmt.Println(myFunc(7))
}

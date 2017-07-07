package main

import "fmt"

/*
Write a function which takes an integer. The function will have two returns.
The first return should be the argument divided by 2.
The second return should be a bool that let’s us know whether or not the argument was even.
For example:
half(1) returns (0, false)
half(2) returns (1, true)
*/

func numF(x int) (float64, bool) {
	var even bool

	if x%2 == 0 {
		even = true
	} else {
		even = false
	}
	y := float64(x)

	return (y / 2), even
}

func main() {
	fmt.Println(numF(2))
	fmt.Println(numF(4))
	fmt.Println(numF(7))
}

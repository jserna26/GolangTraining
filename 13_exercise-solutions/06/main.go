package main

import "fmt"

func main() {
	var multThree, multFive int
	for i := 0; i <= 100; i++ {
		multThree = i % 3
		multFive = i % 5
		fmt.Print(i)
		if (multThree + multFive) == 0 {
			fmt.Print(" FizzBuzz\n")
			continue
		} else if multThree == 0 {
			fmt.Print(" Fizz\n")
		} else if multFive == 0 {
			fmt.Print(" Buzz\n")
		}

	}

}

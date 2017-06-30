package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Print(i)
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(" FizzBuzz \n")
			continue
		} else if i%3 == 0 {
			fmt.Print(" Fizz\n")
		} else if i%5 == 0 {
			fmt.Print(" Buzz\n")
		} else {
			fmt.Println()
		}

	}

}

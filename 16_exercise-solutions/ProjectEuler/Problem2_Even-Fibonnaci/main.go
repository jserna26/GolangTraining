package main

import "fmt"

/**********************************************************
Even Fibonacci numbers
Problem 2

Each new term in the Fibonacci sequence is generated by adding the previous two
terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.

fib
************************************************************/
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}

func sumValList(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	var fibList []int
	var evenFibList []int
	var fibVal int
	var count int

	for {
		fibVal = fib(count)
		if fibVal > 4000000 {
			break
		} else {
			fibList = append(fibList, fibVal)
			if fibVal%2 == 0 {
				evenFibList = append(evenFibList, fibVal)
			}
		}
		count++
	}
	fmt.Println("Sum of fibonnaci values up to 4000000!", sumValList(evenFibList...))

}

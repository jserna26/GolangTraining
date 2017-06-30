package main

import "fmt"

func main() {
	//Zero initiliaze num vars
	var num1, num2 int

	//Ask user for small number
	fmt.Print("Please enter a number: ")
	//Scan first input and assign value to memory address location for num1
	fmt.Scan(&num1)
	fmt.Println("Please enter a number larger than the first: ")
	fmt.Scan(&num2)
	for {
		if num2 > num1 {
			break
		} else {
			fmt.Println("Please try again. Please enter a number larger than the first!")
			fmt.Scan(&num2)
		}
	}
	fmt.Printf("The remainder of %d / %d is: %d", num2, num1, num2%num1)

}

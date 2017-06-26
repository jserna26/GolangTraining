package main

import "fmt"

func main() {
	//Printing memory address of variable stored as well as the decimal value
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address is: ", &a)
	fmt.Printf("ax in decimal is: %d", &a)
}

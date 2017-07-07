package main

import "fmt"

func main() {

	//No name identifier means it is anonymous
	func() {
		fmt.Println("Yup buddy")
	}() //Makes it self executing
}

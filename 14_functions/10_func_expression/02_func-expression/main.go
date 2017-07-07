package main

import "fmt"

func main() {
	//Note when you assign a function to a variable it is known
	//as a func expression
	greeting := func() {
		fmt.Println("Hello world!")
	}
	greeting()
	fmt.Printf("%T", greeting)

}

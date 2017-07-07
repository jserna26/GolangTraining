package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Printf("%T", greet)
	fmt.Println(greet())
}

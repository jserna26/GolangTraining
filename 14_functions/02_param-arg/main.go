package main

import "fmt"

func main() {
	greet("Jon")
	greet("Danielle")
}

func greet(name string) {
	fmt.Println(name)
}

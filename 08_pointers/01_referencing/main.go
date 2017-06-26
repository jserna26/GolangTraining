package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	//pointing to a memory address where an int is stored
	var b *int = &a
	fmt.Println(b)
}

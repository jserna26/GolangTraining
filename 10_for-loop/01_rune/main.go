package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v - %v - %v\n", i, string(i), []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", rune(foo))
}

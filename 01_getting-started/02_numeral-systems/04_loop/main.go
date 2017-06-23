package main

import "fmt"

func main() {
	for i := 500; i < 2000; i++ {
		fmt.Printf("%d \t %b \t %X \n", i, i, i)
	}
}

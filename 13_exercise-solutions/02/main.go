package main

import "fmt"

func main() {
	var x string
	fmt.Print("Type your name: ")
	fmt.Scan(&x)
	fmt.Println("Hello, how's it going ", x)
}

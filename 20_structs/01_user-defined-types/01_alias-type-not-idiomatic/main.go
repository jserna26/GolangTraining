package main

import "fmt"

//Alias for an int example.
//It is a bad practice to alias types!!!
type foo int

func main() {
	var myAge foo
	myAge = 27
	fmt.Println(myAge)

	var yourAge int
	yourAge = 26

	//This does not work. Go is statically typed
	//	fmt.Println(myAge + yourAge)
	fmt.Println(int(myAge) + yourAge)

}

package main

import "fmt"

func main() {
	fmt.Println(greet("Danielle ", "Jon"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

//fmt.Sprint() formats unlimited number of args of any type
//to a concatenated string. It does not print to stdout

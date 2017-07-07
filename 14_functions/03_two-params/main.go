package main

import "fmt"

func main() {
	greet("Jon", "Danielle")
}

func greet(mname, fname string) {
	fmt.Println(mname, "&", fname)
}

//greet declard with two parameters
//Greet is called passing arguments

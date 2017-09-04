package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		"Jonathan",
		"Serna",
		27,
	}

	fmt.Println(p1.first, p1.age)

}

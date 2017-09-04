package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//When you include person as the receive to the function. Any value of the type person can now call this method.
func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{
		"Jonathan",
		"Serna",
		27,
	}

	p2 := person{
		"Danielle",
		"Serna",
		26,
	}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

}

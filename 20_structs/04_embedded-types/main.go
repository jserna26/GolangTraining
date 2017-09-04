package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Jonathan",
			Last:  "Serna",
			Age:   27,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Danielle",
			Last:  "Serna",
			Age:   26,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.LicenseToKill)

}

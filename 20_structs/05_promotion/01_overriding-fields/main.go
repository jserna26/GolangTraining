package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Jonathan",
			Last:  "Serna",
			Age:   27,
		},
		First:         "Double zero 7", //Overridden with First in the outer field
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Danielle", //Overridden with First in the outer field
			Last:  "Serna",
			Age:   26,
		},
		First:         "Double zero 7 wife",
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.LicenseToKill)

}

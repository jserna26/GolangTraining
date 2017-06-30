package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("Did'nt know")
	}

}

func main() {
	SwitchOnType(1231)
	SwitchOnType("ssfsdf")
	t := Contact{
		greeting: "Howdy",
		name:     "Jonathan",
	}
	SwitchOnType(t)

}

package main

import "fmt"

func main() {
	switch "Serna" {
	case "Danielle":
		fmt.Println("Hey Danielle")
	case "Serna":
		fmt.Println("Whaddup Dawg")
		fallthrough
	case "Tobias":
		fmt.Println("Whaddup Toby")
		fallthrough
	case "Jeff":
		fmt.Println("Whaddup Jeff")
	case "Jerome":
		fmt.Println("Whaddup Jerome")
	default:
		fmt.Println("Come on bro! You are defaulting!")

	}

}

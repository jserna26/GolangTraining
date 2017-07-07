package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m["Jon"])
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z map[string]int) {
	z["Jon"] = 44
	fmt.Println(z)
}

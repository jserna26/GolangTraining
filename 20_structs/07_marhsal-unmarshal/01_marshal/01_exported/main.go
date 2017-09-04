package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string `json:"-"`
	Age         int    `json:"wisdom-score"`
	notExported int
}

func main() {
	p1 := Person{
		"Jonathan",
		"Serna",
		27,
		007,
	}

	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

}

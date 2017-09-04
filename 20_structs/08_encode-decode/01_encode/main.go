package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{"Jonathan", "Serna", 27}

	json.NewEncoder(os.Stdout).Encode(p1)

}

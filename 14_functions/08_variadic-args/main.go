package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 58}
	//data... pull items out of slise and pass it one at a time
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

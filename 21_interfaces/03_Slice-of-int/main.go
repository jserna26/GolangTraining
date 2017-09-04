package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 1, 4, 5, 6, 2, 89, 10, 56, 55}
	fmt.Printf("Slice unsorted: %v\n", n)
	sort.Sort(sort.IntSlice(n))
	fmt.Printf("Slice sorted: %v", n)

}

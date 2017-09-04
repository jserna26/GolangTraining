package main

import (
	"fmt"
	"sort"
)

/*
1) type people[]string
studygroup {"Zeno", "john", "Al", "jenny"}*/
type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	p2 := []string{"Jackson", "Zeno", "John", "Al", "Jenny"}
	fmt.Println(p2)
	i := sort.Reverse(sort.StringSlice(p2))
	fmt.Println(i)

}

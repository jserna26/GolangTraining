package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Get the Moby Dick Book
	res, err := http.Get("https://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//Scan the page
	//NewScanner takes a reader and res.body implements the reader interface
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//create slice to hold the counts
	buckets := make([]int, 200)
	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:122])

}

func HashBucket(word string) int {
	return int(word[0])
}

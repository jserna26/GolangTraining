package main

import (
	"fmt"
	"time"
)

func main() {
	//Unbuffered Channel. No limit into how many types can be pushed to the Channel
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//Once a value is placed on the channel. This is blocked until a reciever takes the value off the channel.
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(time.Second)

}

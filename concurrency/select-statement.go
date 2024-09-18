package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "result from ch1"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "result from ch2"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}

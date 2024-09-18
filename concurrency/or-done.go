package main

import (
	"fmt"
)

func orDone(done chan struct{}, in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				fmt.Printf("go-routine cancelled\n")
			case val, ok := <-in:
				if !ok {
					return
				}
				out <- val
			}
		}
	}()
	return out
}

/*
func main() {
	done := make(chan struct{})
	in := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
			fmt.Printf("Input :%d sent\n", i)
		}
		close(in)
	}()

	go func() {
		time.Sleep(time.Second)
		done <- struct{}{}
		close(done)
		fmt.Println("Sent signal to cancell go-routine")
	}()

	result := orDone(done, in)
	for r := range result {
		fmt.Printf("result val:%d\n", r)
	}
}*/

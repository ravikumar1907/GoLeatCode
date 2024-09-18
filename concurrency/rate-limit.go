package main

import (
	"fmt"
	"time"
)

func rateLimit() {
	rate := time.Second
	limiter := time.NewTicker(rate)
	defer limiter.Stop()
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	for req := range requests {
		fmt.Println("Start handling request")
		<-limiter.C
		fmt.Println("Handling request", req)
	}
}

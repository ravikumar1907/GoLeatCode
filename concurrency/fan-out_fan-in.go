package main

import (
	"fmt"
	"sync"
)

func fanOut(data []int) chan int {
	out := make(chan int, len(data))
	go func() {
		var wg sync.WaitGroup
		for i, v := range data {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				fmt.Printf("Task :%d, data:%d\n", i+1, val)
				out <- val * 2
			}(v)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func fanIn(in <-chan int) []int {
	var result []int
	for v := range in {
		fmt.Printf("Processed data :%d\n", v)
		result = append(result, v)
	}
	return result
}

// func main() {
// 	data := []int{1, 2, 3, 4, 5}
// 	in := fanOut(data)
// 	result := fanIn(in)
// 	fmt.Printf("Fainal result:%v", result)

// }

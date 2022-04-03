package main

import (
	"fmt"
	"sync"
)

func PrintOddEven() {
	odd := make(chan int)
	even := make(chan int)
	data := make(chan int)
	go func() {
		for i := 1; i < 100; i = i + 2 {
			data <- i
			even <- 1
			<-odd
		}
	}()
	go func() {
		for i := 2; i <= 100; i = i + 2 {
			<-even
			data <- i
			odd <- 1
		}
		close(data)
	}()
	for {
		v, ok := <-data
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

}
func Permute(a []rune, i, n int) {
	if i >= n {
		fmt.Println(string(a))
	}
	for j := i; j < n; j++ {
		a[i], a[j] = a[j], a[i]
		Permute(a, i+1, n)
		a[i], a[j] = a[j], a[i]
	}
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		} else if len(right) == 0 {
			return append(result, left...)
		} else if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func SqaureANumber(c, quit chan int) {
	i := 1
	for {
		select {
		case c <- i * i:
			i++
		case <-quit:
			return
		}
	}
}

func SumOfSqaures(n int) {

	read := make(chan int)
	done := make(chan int)
	sum := 0
	go SqaureANumber(read, done)
	for i := 1; i <= n; i++ {
		sum += <-read
	}
	done <- 1
	fmt.Println(sum)
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	//	done := make(chan bool)
	var left []int
	var right []int
	m := len(a) / 2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		left = mergeSort(a[:m])
		wg.Done()
	}()
	go func() {
		right = mergeSort(a[m:])
		wg.Done()
	}()
	wg.Wait()
	return merge(left, right)
}

func generate(n int) chan int {
	numbers := make(chan int)
	go func(num int) {
		for i := 1; i <= n; i++ {
			numbers <- i
		}
		close(numbers)
	}(n)
	return numbers
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func FanIn(chan1, chan2 <-chan int) chan int {
	result := make(chan int)
	go func() {
		for {
			select {
			case v := <-chan1:
				result <- v
			case v := <-chan2:
				result <- v
			}
		}
	}()
	return result
}

/*
func main() {
	//PrintOddEven()
	//Permute([]rune("abc"), 0, 3)
	n := 10
	SumOfSqaures(n)
	fmt.Println(mergeSort([]int{3, 7, 1, 6, 2, 5, 9, 10, 4, 8}))
}*/

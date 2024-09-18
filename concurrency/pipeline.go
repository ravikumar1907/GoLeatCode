package main

/*
Pipeline: A series of stages connected by channels where each stage performs some processing on data and passes it to the next stage.
Each stage is typically a goroutine, and channels are used to pass data between stages.
Use Case: Data processing tasks where each step is dependent on the previous one, like transforming, filtering, or aggregating data.
go

*/

func stage1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i + 1
		}
		close(out)
	}()
	return out
}

/*func main() {
	input := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Passing through pipeline stages
	s1 := stage1(input)
	s2 := stage2(s1)

	for result := range s2 {
		fmt.Println(result)
	}
}*/

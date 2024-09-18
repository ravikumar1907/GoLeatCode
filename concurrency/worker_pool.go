package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wId int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker :%d started job:%d\n", wId, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker :%d finished job:%d\n", wId, job)
		results <- job * 2
	}
}

// func main() {
// 	numWorkers := 3
// 	numJobs := 5
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)
// 	var wg sync.WaitGroup

// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}
// 	for i := 1; i <= numJobs; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)
// 	/* Good to have below */
// 	// go func() {
// 	// 	wg.Wait()
// 	// 	close(results)
// 	// }()
// 	wg.Wait()
// 	close(results)
// 	for res := range results {
// 		fmt.Printf("result :%d\n", res)
// 	}
// }

/*
The pattern used in the updated solution is commonly referred to as the "fan-out, fan-in" pattern in Go concurrency.

Explanation of the Pattern
Fan-Out:

The "fan-out" part of the pattern refers to starting multiple goroutines (workers) to handle parts of a workload concurrently.
In the example, multiple worker goroutines (worker function) are started, each listening to a shared jobs channel and processing jobs concurrently.
Fan-In:

The "fan-in" part refers to aggregating the results from multiple goroutines into a single channel for further processing.
Here, all workers send their results to the shared results channel, and the main goroutine reads from this channel to collect and handle the results.

Additional Pattern Used: Goroutine with WaitGroup
Another pattern demonstrated is the use of a goroutine with a WaitGroup to wait for multiple goroutines to finish their work:

The anonymous goroutine (go func() { ... }()) waits for all workers to complete (wg.Wait()) and then safely closes the results channel.
This pattern ensures that the main function does not block indefinitely and that resources (like channels) are properly managed.

Why Use This Pattern?
Efficient Concurrency: Enables concurrent processing, maximizing the use of system resources.
Safe Synchronization: Ensures safe closure of channels and synchronization among multiple goroutines.
Scalable: Can easily scale the number of workers to match the workload or system capabilities.

Summary
Fan-Out: Start multiple worker goroutines to perform concurrent tasks.
Fan-In: Aggregate results from multiple goroutines into a single channel for processing.
Goroutine with WaitGroup: Wait for completion and ensure proper resource cleanup.

Aspect				Worker Pool Pattern														Fan-Out, Fan-In Pattern
Purpose				To limit the number of goroutines processing tasks concurrently.		To process tasks concurrently and aggregate results from multiple sources.
Structure			Fixed number of worker goroutines pull tasks from a shared job queue.	Multiple goroutines (fan-out) process tasks and send their results to a single channel (fan-in).
Concurrency Control	Controls the level of concurrency by limiting the number of workers.	Maximizes concurrency without directly controlling the number of goroutines.
Resource Management	Designed to manage and limit resource usage (e.g., CPU, memory).		Designed to maximize throughput, with less emphasis on resource constraints.
Use Case			Suitable for limiting the rate of processing, managing 					Suitable for processing a large number of tasks in parallel and aggregating results quickly.
					resource-intensive tasks, or preventing overload.
Typical Flow		- Tasks are added to a queue.											- Tasks are distributed to multiple goroutines (fan-out).
					- Workers (goroutines) pull tasks from the queue.						- Each goroutine processes a task and sends results to a common channel (fan-in).
					- Workers process tasks and push results to a results channel
					  or another handler.

*/

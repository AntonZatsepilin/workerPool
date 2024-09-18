package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	const jobsCount, workerCount = 40, 20

	jobs := make(chan int, jobsCount)
	results := make(chan int, workerCount)

	for i := 0; i < workerCount; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}
	defer close(jobs)

	for i := 0; i < jobsCount; i++ {
		fmt.Printf("result [%d] : value = %d\n", i+1, <-results)
	}
	fmt.Println("TIME ELAPSED:", time.Since(t).String())
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("worker [%d] finished\n", id)
		results <- j * j
	}
}

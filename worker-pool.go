package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finish job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 0; a < numJobs; a++ {
		<-result
	}
}

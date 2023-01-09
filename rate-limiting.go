package main

import (
	"fmt"
	"time"
)

func main() {

	// Default regular limiter
	resquest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		resquest <- i
	}
	close(resquest)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range resquest {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Burst limiter
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

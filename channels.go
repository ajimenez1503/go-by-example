package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// Channel unbuffered
	messages := make(chan string)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// Channel buffered
	messagesBuffered := make(chan string, 2)
	go func() {
		messagesBuffered <- "ping1"
		messagesBuffered <- "ping2"
	}()

	fmt.Println(<-messagesBuffered)
	fmt.Println(<-messagesBuffered)

	// Channel sysncrhonization
	done := make(chan bool, 1)
	go worker(done)

	<-done

	// Channel direccion
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

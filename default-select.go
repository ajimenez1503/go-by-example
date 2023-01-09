package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	case sig := <-signals:
		fmt.Println(sig)
	default:
		fmt.Println("not activity")
	}

	select {
	case messages <- "hi":
		fmt.Println("Message sent")
	default:
		fmt.Println("no message sent")
	}
}

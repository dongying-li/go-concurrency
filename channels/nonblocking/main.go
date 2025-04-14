package main

import (
	"time"
	"fmt"
)

func main() {
	data := make(chan string)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Session ended")
			case msg := <-data:
				fmt.Println(msg)
			default:
				fmt.Println("Waiting")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	numMessages := 5
	for i := range numMessages {
		data <- "Message " + string(rune('1' + i))
		time.Sleep(1 * time.Second)
	}

	done <- true
}

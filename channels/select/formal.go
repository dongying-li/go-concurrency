package main

import (
	"fmt"
)

func formal() {
	msgStream := make(chan string)

	go func(msgStream chan<- string) {
		for range 3 {
			msgStream <- "some messages..."
		}
		close(msgStream)
	}(msgStream)

	for {
		msg, ok := <-msgStream
		if !ok {
			fmt.Println("channel closed...")
			break
		}
		fmt.Println(msg)
	}
}

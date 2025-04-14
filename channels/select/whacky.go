package main

import (
	"fmt"
	"time"
)

// Guess the possible output of the following program

func whacky() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		close(ch2)
	}()

	time.Sleep(2 * time.Second)
	for range 3 {
		select {
		case msg, ok := <-ch1:
			fmt.Println("ch1", msg, ok)
		case msg, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 has been closed")
				continue
			}
			fmt.Println("ch2", msg, ok)
		}
	}

	msg, ok := <-ch1
	fmt.Println(msg, ok) // 1 true
	msg, ok = <-ch2
	fmt.Println(msg, ok)

	fmt.Println("End of the program")
}

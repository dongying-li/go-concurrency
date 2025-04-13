package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	
	// make ch reach its capacity
	ch <- 1
	ch <- 2

	go func() {
		time.Sleep(2 * time.Second)
		// A whacky behaviour of this program is that the following print is not guaranteed
		// Why? code evaluated from right to left. After <-ch is done, line 24 will no longer be blocked,
		// 		and the program will move on and end even before the print is done.
		fmt.Printf("Received: %d\n", <-ch)
	}()

	fmt.Println("Blocking")
	ch <- 3
	fmt.Println("Unblocked")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	
	fmt.Println("=========== Ordered Sync ===========")
	orderedSync()

	fmt.Println("=========== Unordered Sync ===========")
	unorderedSync()
}

func orderedSync() {

	msgCh := make(chan string)
	numMessages := 3

	go func(numMessages int) {
		for i := range numMessages {
			msgCh <- "Message " + string(rune('1' + i))
		}
		
		close(msgCh)
	}(numMessages)

	for msg := range msgCh {
		fmt.Println(msg)
	}
}

func unorderedSync() {
	done := make(chan int, 3)

	numGoroutines := 3

	for i := range numGoroutines {
		go func(i int) {
			fmt.Printf("goroutine %d\n", i + 1)
			time.Sleep(1 * time.Second)
			done <- i + 1
		}(i)
	}

	// may print 3 goroutines or just 2 if the second for loop runs < numGoroutines iterations
	// for range numGoroutines - 1 {
	// 	<-done
	// }

	for range numGoroutines {
		<-done
	}
}

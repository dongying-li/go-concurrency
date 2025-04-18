package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	done := make(chan bool)
	timer := time.Now().Add(5000 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, timer)
	defer cancelCtx()
	
	printCh := make(chan int)
	go doAnother(ctx, printCh, done)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1500 * time.Millisecond)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	<-done

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int, done chan<- bool) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			done <- true
			return
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}


func main() {
	ctx := context.Background()
	doSomething(ctx)
}

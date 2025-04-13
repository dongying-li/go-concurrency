package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var err error

	go func() {
		err = doWork()
	}()
	
	// Uncommend the line below to avoid goroutine leak
	// time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Completed successfully")
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return errors.New("Some errors")
}

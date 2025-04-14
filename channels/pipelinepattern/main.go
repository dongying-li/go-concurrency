package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go process(ch1)
	go filter(ch1, ch2)

	for data := range ch2 {
		fmt.Println(data)
	}

	fmt.Println("End of the pipeline")
}

func process(out chan<- int) {
	for i := range 10 {
		out <- i
	}
	close(out)
}

func filter(in <-chan int, out chan<- int) {
	for data := range in {
		if data % 2 == 0 {
			out <- data
		}
	}
	close(out)
}

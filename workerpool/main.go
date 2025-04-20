package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerID int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for jobID := range jobs {
		fmt.Printf("Worker %d received job %d\n", workerID, jobID)
		time.Sleep(time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", workerID, jobID)
	}
}

func main() {
	var wg sync.WaitGroup
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for i := range 3{
		go worker(i, jobs, results, &wg)
	}

	wg.Add(3)
	go func() {
		wg.Wait()
		close(results)
	}()

	for jobID := range numJobs {
		jobs <- jobID
	}
	close(jobs)

	for res := range results {
		fmt.Println(res)
	}

}

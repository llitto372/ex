package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range jobs {
		fmt.Printf("Воркер %d, работа номер %d\n", id, val)
	}

}

func main() {
	const (
		numJobs    = 20
		numWorkers = 5
	)

	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}

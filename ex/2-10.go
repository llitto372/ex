package main

import (
	"fmt"
	"sync"
)

func work(id int, jobsCh chan int, resultCh chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := make([]int, 1)
	tempRes := 0
	for val := range jobsCh {
		tempRes = val * 2
		res = append(res, tempRes)
	}
	resultCh <- res

}

func main() {
	const (
		jobsCount    = 20
		workersCount = 5
	)
	jobsCh := make(chan int)
	resultCh := make(chan []int, workersCount)
	var wg sync.WaitGroup
	for i := 1; i <= workersCount; i++ {
		wg.Add(1)
		go work(i, jobsCh, resultCh, &wg)
	}
	for j := 1; j <= jobsCount; j++ {
		jobsCh <- j
	}
	close(jobsCh)

	wg.Wait()

	close(resultCh)

	// for value := range resultCh {
	// 	for _, val := range value {
	// 		fmt.Println(val)
	// 	}
	// }
	fmt.Println(<-resultCh)
}

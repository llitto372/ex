package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func SumPart(nums []int, resultCh chan int) {
	defer wg.Done()
	res := 0
	for _, val := range nums {
		res += val
	}
	resultCh <- res
}

func main() {
	const parts = 10
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}
	resultCh := make(chan int, parts)

	segm := len(nums) / 10
	for i := 0; i < parts; i++ {
		huy := nums[segm*i : segm*(i+1)]
		wg.Add(1)
		go SumPart(huy, resultCh)
	}
	wg.Wait()
	close(resultCh)
	sumCh := 0

	for val := range resultCh {
		sumCh += val
	}
	fmt.Println("Результат: ", sumCh)
}

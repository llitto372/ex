package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func genenerator(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(5) + 1
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go genenerator(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}

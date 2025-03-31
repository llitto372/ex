package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func producer(ch chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(5) + 1
	}
	close(ch)
}

func square(ch1 chan int, ch2 chan int) {
	defer wg.Done()
	for val := range ch1 {
		ch2 <- val * val
	}
	close(ch2)
}

func print(ch2 chan int) {
	defer wg.Done()
	for val := range ch2 {
		fmt.Println("Результат: ", val)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(3)
	go producer(ch1)
	go square(ch1, ch2)
	go print(ch2)
	wg.Wait()
}

// задача почти идентичная прошлой, без затупов по мат части, только я при инициализации
// каналов каким то хуем в ch написал русскую с, теперь понял, что русские символы подсвечивает
// желтым

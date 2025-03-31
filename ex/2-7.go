package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type SafeNumber struct {
	value int
	mu    sync.RWMutex
}

func (s *SafeNumber) Get() int {
	defer wg.Done()
	s.mu.RLock()
	fmt.Println("Get", s.value)
	defer s.mu.RUnlock()
	return s.value
}

func (s *SafeNumber) Set(v int) {
	defer wg.Done()
	for i := 0; i < v; i++ {
		s.mu.Lock()
		fmt.Println("Set")
		s.value++
		s.mu.Unlock()
	}

}
func main() {
	num := SafeNumber{value: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go num.Get()
	}
	wg.Add(1)
	go num.Set(5)

	wg.Wait()
}

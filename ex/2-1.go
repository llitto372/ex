package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
			fmt.Println(counter)
		}()

	}
	wg.Wait()
	fmt.Println("Result:", counter)
}

//без затупов, базовая задачка на понимание работы мьютексов и веитгруп
//пощупал без вейтгруп не корректно очевидно работает, потому что не ждет
//окончания всех горутин, а завершается после конца main

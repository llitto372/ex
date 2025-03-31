package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func count(val *int32) {
	atomic.AddInt32(val, 1)
}

func main() {
	var val int32 = 5
	for i := 0; i < 10; i++ {
		go count(&val)
		go count(&val)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(val)
}

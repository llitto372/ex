package main

import (
	"context"
	"fmt"
)

func doSmth(ctx context.Context, times int, resCh chan<- int) {
	count := 0
	for i := 0; i < times; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			count++

		}
	}
	select {
	case <-ctx.Done():
		return
	case resCh <- count:
	}

}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	resCh := make(chan int)
	go doSmth(ctx, 500, resCh)
	go doSmth(ctx, 1001, resCh)
	fmt.Println(<-resCh)
	cancel()
}

// один сплошной затуп, царь затупов в деле

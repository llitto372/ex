package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Конец")
			return
		default:
			fmt.Println("tick")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(4 * time.Second)
}

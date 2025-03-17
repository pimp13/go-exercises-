package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx)

	time.Sleep(2 * time.Second)
}

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Microsecond):
		fmt.Println("Work Done!")
	case <-ctx.Done():
		fmt.Println("Canceled:", ctx.Err())
	}
}

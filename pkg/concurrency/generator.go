package concurrency

import (
	"context"
	"fmt"
	"time"
)

func generator[T any](context context.Context, list []T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range list {
			select {
			case ch <- v:
				fmt.Println("NOW LOADING...")
				time.Sleep(1 * time.Second)
			case <-context.Done():
				fmt.Println("ABORT...")
				return
			}
		}
	}()
	return ch
}

func RunGenerator() {
	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	for item := range generator(context, []int{1, 2, 3, 4, 5}) {
		if item == 4 {
			cancel()
		}
		fmt.Println(item)
	}
	fmt.Println("EXIT...")
}

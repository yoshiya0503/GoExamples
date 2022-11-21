package concurrency

import (
	"context"
	"fmt"
	"time"
)

func RunForSelect() {
	context, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		defer close(ch1)
		time.Sleep(10 * time.Second)
		ch1 <- "call ch1 func"
	}()

	go func() {
		defer close(ch2)
		ch2 <- "call ch2 func"
	}()

	for {
		select {
		case result, ok := <-ch1:
			if ok {
				fmt.Println(result)
			}
		case result, ok := <-ch2:
			if ok {
				fmt.Println(result)
			}
		case <-context.Done():
			fmt.Println("canceled...")
			return
		}
	}
}

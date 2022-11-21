package concurrency

import (
	"context"
	"fmt"
	"time"
)

type Pipeline[T any] func() <-chan T

func NewPipeline[T any](context context.Context, args ...T) Pipeline[T] {
	return Pipeline[T](func() <-chan T {
		stream := make(chan T)
		go func() {
			defer close(stream)
			for _, arg := range args {
				select {
				case stream <- arg:
				case <-context.Done():
					return
				}
			}
		}()
		return stream
	})
}

func (p Pipeline[T]) Add(context context.Context, callback func(T) T) Pipeline[T] {
	return Pipeline[T](func() <-chan T {
		stream := make(chan T)
		go func() {
			defer close(stream)
			for arg := range p() {
				select {
				case stream <- callback(arg):
				case <-context.Done():
					return
				}

			}
		}()
		return stream
	})
}

func RunPipeline() {
	fetchWebService := func(html string) string {
		time.Sleep(3 * time.Second)
		return fmt.Sprintf("<html><title>%s</title></html>", html)
	}
	consumeHTML := func(html string) string {
		time.Sleep(1 * time.Second)
		return fmt.Sprintf("加工されたHTMLはこちら... <html><title>%s</title></html>", html)
	}
	writeCSV := func(consumedHTML string) string {
		fmt.Println("[writeCSV] CSVを書き出しています")
		time.Sleep(1 * time.Second)
		return consumedHTML
	}
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream := NewPipeline(context, "https://test.com", "https://github.com", "https://youtube.com").
		Add(context, fetchWebService).
		Add(context, consumeHTML).
		Add(context, writeCSV)
	for r := range stream() {
		fmt.Println(r)
	}
}

package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	// recv chan
	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()

	// send chan
	go func() {
		defer wg.Done()
		defer close(c) //发送完成后关闭通道，通道关闭后依然可以接受已缓冲的数据

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()
}

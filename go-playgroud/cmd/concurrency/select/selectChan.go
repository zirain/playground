package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	// recv
	go func() {
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)

			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if !ok {
				return
			}

			println(name, x)
		}
	}()

	// send
	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	wg.Wait()
}

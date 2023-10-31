package main

import(
	"time"
	"sync"
)

//可以再多处使用wait阻塞
func main()  {
	var wg sync.WaitGroup

	wg.Add(1)

	go func ()  {
		wg.Wait()
		println("wait exit.")
	}()

	go func ()  {
		time.Sleep(time.Second)
		println("done.")
		wg.Done()
	}()

	wg.Wait()
	println("main exit.")
}
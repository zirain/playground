package main

import(
	"time"
)

//使用通道（channel）阻塞进程，等待并发任务结束
func main()  {
	exit :=make(chan struct{})

	go func ()  {
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit)
	}()

	println("main...")
	<-exit
	println("main exit.")
}
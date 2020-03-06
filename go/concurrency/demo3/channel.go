package main

import(

)

func main()  {
	done := make(chan struct{})
	c := make(chan string)

	go func ()  {
		s := <- c
		println(s)
		close(done)
	}()

	c <- "hi!"

	<- done
}
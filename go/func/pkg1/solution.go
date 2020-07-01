package main

import "fmt"

type A func(int, int)

func (f A) Serve() {
	fmt.Println("serve2")
}

func serve(int, int) {
	fmt.Println("serve1")
}

func main() {
	a := A(serve)
	a(1, 2)
	a.Serve()
}

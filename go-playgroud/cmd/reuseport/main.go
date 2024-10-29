package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	reuseport "github.com/kavu/go_reuseport"
)

func main() {
	// listener, err := reuseport.Listen("tcp", ":8881")
	listener, err := reuseport.Listen("tcp6", ":8881")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server started on %s\n", listener.Addr().String())
	defer listener.Close()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(os.Getgid())
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	panic(server.Serve(listener))
}

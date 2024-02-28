package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed jwks.json
var jwks string

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(jwks))
	})

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("start to listen :8080")
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("fail to listen: ", err)
	}
}

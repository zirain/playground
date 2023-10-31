package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	err = &os.PathError{
		Path: "err path",
		Op:   "err op",
		Err:  errors.New("path err"),
	}
	// 无法确定wrappedErr的类型
	wrappedErr := fmt.Errorf("wrappedErr %w", err)
	if _, ok := wrappedErr.(*os.PathError); ok {
		fmt.Println("assert wrappedErr  os.PathError")
	}
	// 用As来断言wrappedErr的类型
	var pathErr *os.PathError
	if errors.As(wrappedErr, &pathErr) {
		fmt.Println("as wrappedErr os.PathError")
	}

	fmt.Printf("err: %v\n", wrappedErr)
}
package main

import (
	"fmt"
)

// panic仅有最后一个可以被revover捕获。
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}

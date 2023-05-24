package main

import "fmt"

// 先执行 defer中的函数，后抛出error
func main() {
	defer func() {
		fmt.Println("exec defer")
	}()

	panic("panic error")
}

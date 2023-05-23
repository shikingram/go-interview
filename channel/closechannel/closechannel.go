package main

import "fmt"

func main() {
	ch := make(chan struct{})

	go func() {
		defer close(ch)
		ch <- struct{}{}
	}()

	// 已经关闭的channel是可以继续读取数据的，但是不能写入
	a := <-ch
	fmt.Println(a)
}

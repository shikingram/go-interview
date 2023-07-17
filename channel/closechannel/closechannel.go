package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
		wg.Done()
	}()

	// 已经关闭的channel是可以继续读取数据的，但是不能写入
	wg.Wait()
	for i := 0; i < 7; i++ {
		a := <-ch
		fmt.Println(a)
	}
	ch <- 6 //panic
}

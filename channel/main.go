package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})
	// 此处重复关闭会导致panic
	// panic: close of closed channel
	defer func() {
		close(ch)
		fmt.Println("main defer over")
	}()

	go func() {
		defer func() {
			close(ch)
			fmt.Println("gorutine defer over")
		}()
		ch <- struct{}{}
		fmt.Println("go rutine over")
	}()
	i := 0
	fmt.Println("waiting")

	// NOTE: 遍历时如果管道没有关闭，会deadlock
	// fatal error: all goroutines are asleep - deadlock!
	for range ch {
		i++
	}
	fmt.Printf("%d \n", i)
}

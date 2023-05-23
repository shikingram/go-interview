package main

import (
	"sync"
	"sync/atomic"
)

// fatal error: all goroutines are asleep - deadlock!
// 此处死锁，因为waitgroup 使用了值传递，wg.Done()只是关闭了复制的wg变量
func main() {

	var wg sync.WaitGroup

	ans := int64(0)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go newGoRoutine(wg, &ans)
	}
	wg.Wait()
}

func newGoRoutine(wg sync.WaitGroup, i *int64) {
	defer wg.Done()
	atomic.AddInt64(i, 1)
}

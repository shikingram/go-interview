package main

import (
	"fmt"
	"interview/rwmutex/utils"
	"sync"
	"time"
)

func main() {
	// 对并发读操作进行测试
	m := utils.NewMyConCurrentMap()

	wg1 := sync.WaitGroup{}
	m.Put("key1", "val1")
	m.Put("key2", "val2")
	wg1.Add(2)
	for i := 1; i < 3; i++ {
		go func(i int, wg1 *sync.WaitGroup) {
			_, _ = m.Get(fmt.Sprintf("key%d", i))
			// 此处输出的数据可能是先输出val1 也可能是先输出val2 因为是并发读  没有先后顺序
			wg1.Done()
		}(i, &wg1)
	}
	wg1.Wait()

	m.Delete("key1")
	m.Delete("key2")

	wg := sync.WaitGroup{}
	// 对并发读+写的场景进行测试
	wg.Add(3)
	for i := 3; i < 6; i++ {
		go func(i int, wg *sync.WaitGroup) {
			key := "samekey"
			m.Put(key, time.Now().Format(time.RFC3339Nano))
			time.Sleep(time.Microsecond * 30) //放大实现现象
			_, _ = m.Get(key)                 // 上面进行了睡眠操作，所有gorutine都进行了samekey写入，此处所有gorutine的读操作都是读取了最后一次写入的val
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

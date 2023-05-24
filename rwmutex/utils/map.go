package utils

import (
	"fmt"
	"sync"
)

/*
读写锁和互斥锁的区别：
并发对临界资源的操作广义上可以归类为以下三种类型
1. 并发读
2. 并发写
3. 并发读+写

读写锁在处理  并发读的场景时，可以允许多个gorutine进行读操作
而在处理并发写  并发读+写的场景时，不允许多个gorutine同时进行操作

互斥锁对三种场景统一进行加锁不允许同时操作
*/
type MyConCurrentMap struct {
	sync.RWMutex
	mp map[string]string
}

func NewMyConCurrentMap() *MyConCurrentMap {
	return &MyConCurrentMap{
		mp: make(map[string]string),
	}
}

func (m *MyConCurrentMap) Get(key string) (val string, ok bool) {
	m.RWMutex.RLock()
	val, ok = m.mp[key]
	fmt.Println("get", key, "return", val, ok)
	m.RWMutex.RUnlock()
	return val, ok
}

func (m *MyConCurrentMap) Put(key, val string) {
	m.RWMutex.Lock()
	m.mp[key] = val
	fmt.Println("put[", key, "]success")
	m.RWMutex.Unlock()
}

func (m *MyConCurrentMap) Delete(key string) {
	m.RWMutex.Lock()
	delete(m.mp, key)
	m.RWMutex.Unlock()
}

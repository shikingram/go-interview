package main

import (
	"encoding/json"
	"log"
)

type S struct {
	A int
	B *int
	C float64
	d func() string
	e chan struct{}
}

// func channel类型是不能被当做变量序列化的，但是这里是小写，json包不可见，因此不会发生序列化错误
func main() {
	s := S{
		A: 1,
		B: nil,
		C: 12.15,
		d: func() string {
			return "NowCoder"
		},
		e: make(chan struct{}),
	}

	_, err := json.Marshal(s)
	if err != nil {
		log.Printf("err occurred..")
		return
	}
	log.Printf("everything is ok.")
}

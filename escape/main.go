package main

import "fmt"

/*
几种变量逃逸的情况
*/

// 1 指针类型的逃逸
func escape1() *int {
	var a int = 1
	return &a //main.go:7:6: moved to heap: a
}

// 2 栈空间不足
func escape2() {
	s := make([]int, 0, 10000) //make([]int, 0, 10000) escapes to heap
	for index := range s {
		s[index] = index
	}
}

// 3 变量大小不确定
func escape3() {
	number := 10
	s := make([]int, number) //make([]int, number) escapes to heap
	for index := range s {
		s[index] = index
	}
}

// 4 变量类型是动态类型 interface
func escape4() {
	fmt.Println("111") // main.go:29:14: "111" escapes to heap
}

func main() {
	escape1()
	escape2()
	escape3()
	escape4()
}

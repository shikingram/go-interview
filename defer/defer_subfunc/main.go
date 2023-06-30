package main

import "fmt"

func function(index int, value int) int {

	fmt.Println(index)

	return index
}

func main() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}

/**
- defer压栈function1，压栈函数地址、形参1、形参2(调用function3) --> 打印3
- defer压栈function2，压栈函数地址、形参1、形参2(调用function4) --> 打印4
- defer出栈function2, 调用function2 --> 打印2
- defer出栈function1, 调用function1--> 打印1
**/

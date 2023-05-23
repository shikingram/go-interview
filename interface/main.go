package main

type InterfaceA interface {
	Add(a, b int) int
}

type InterfaceB interface {
	Add(a, b int) int
	Minus(a, b int) int
}

func main() {
	var a InterfaceA
	var b InterfaceB
	// a 中的函数是b种的子集，所以b可以赋值给a,但是a赋值给b会报错
	a = b
	a.Add(1, 2)
}

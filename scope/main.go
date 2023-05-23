package main

import "fmt"

// 这里输出12  因为重新定义了a 作用域就是在{}内
func main() {

	a := 12
	{

		a := 13
		_ = a // make compiler happy
	}

	fmt.Println(a)

}

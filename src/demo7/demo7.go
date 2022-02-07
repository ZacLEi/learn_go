package main

import "fmt"

/**
panic recover

Golang 异常的抛出与捕获，依赖两个内置函数：

panic：抛出异常，使程序崩溃

recover：捕获异常，恢复程序或做收尾工作


*/
func main() {
	//panic("crash")

	setData(20)
	fmt.Println("everything is ok!")
}

func setData(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var arr [10]int
	arr[x] = 88
}

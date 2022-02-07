package main

import "fmt"

/**
主要介绍流程控制的语句
*/
func main() {
	/**
	1,if else
	2.switch-case
	3,for range
	4.goto
	5.defer 延迟执行
	*/
	loopSwitch()
	loopFor()
	string2 := "123"
	defer deferFunc(string2)
	string2 = "456"
	println(string2)

}

func loopFor() {
	a := 1
	for a <= 5 {
		println(a)
		a++
	}

	println("---------------------------")
	for i := 1; i <= 5; i++ {
		println(i)
	}
	println("---------------------------")
	for {
		break
	}
}

func loopSwitch() {
	education := "本科"
	switch education {
	case "博士":
		fmt.Println("我是博士")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
		//fallthrough 穿透功能 不判断下一个条件直接执行
		fallthrough
	case "大专":
		fmt.Println("我是大专生")
	case "高中":
		fmt.Println("我是高中生")
	default:
		fmt.Println("学历未达标..")
	}
}

func deferFunc(string2 string) {
	// defer 会使用调用defer函数时传入的参数快照，如果后续对参数进行了更改，也使用传入的参数
	// defer 是反序调用，类似栈
	// defer 一般用于释放资源
	println("defer.........", string2)
}

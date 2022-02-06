package main

import "fmt"

/*
demo1 主要介绍go的几种创建变量的方法和new函数的使用方法
*/
func main() {
	var string1 string = "asdjlkasjdlasd"
	fmt.Println(string1)

	var str2 = "string2"
	var float1 = 1.01
	fmt.Println(str2, float1)

	//var (
	//	str3 string
	//	num1 int
	//	str4 string
	//)

	// 推荐用这种方式
	name := "fuck you"
	fmt.Println(name)

	// 推荐用这种方式
	name1, name2 := "123", "456"
	fmt.Println(name1, name2)

	// new 函数返回一个指针变量
	ptr := new(int)
	fmt.Println(ptr)
	fmt.Println(*ptr)

}

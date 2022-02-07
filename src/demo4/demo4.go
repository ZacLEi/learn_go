package main

/**
介绍指针的相关内容
*/
func main() {
	str := "指针的相关内容"
	ptr := &str
	println(*ptr)
	//
	astr := new(string)
	*astr = "alskdjlajd"
	//
	//aint := 1
	//var bint *int
	//bint = &aint

	aint := 1
	ptr2 := &aint
	println("普通变量存储的是：", aint)
	println("普通变量存储的是：", *ptr2)
	println("指针变量存储的是：", &aint)
	println("指针变量存储的是：", ptr2)

	// 指针声明后，如果没有初始化，其零值是nil
	// 如果想通过一个函数改变一个数组的值，最好使用切片的方式，因为切片也是引用类型

}

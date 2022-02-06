package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	var arr2 [3]int = [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}

	fmt.Println(arr2)
	fmt.Println(arr3)

	// 表示创建一个长度为4的数组，下标为2的值为3
	arr4 := [4]int{2: 3}
	fmt.Println(arr4)

	//---------------------切片---------------------
	fmt.Println("----------------------------------------")
	myarr := [8]int{0, 1, 2, 3, 4, 5, 6, 7} // 创建一个数组
	// 表示从1开始，到下标为2（3-1）的元素引用
	slice1 := myarr[1:4]
	slice2 := myarr[1:4:4]
	fmt.Printf("%d，%d\n", len(slice1), cap(slice1))
	fmt.Printf("%d, %d\n", len(slice2), cap(slice2))
	// 切片的第三个参数，表示容量到这个下标处，不能小于第二个参数，不能大于底层数组的容量

	// 使用make方法创建切片
	slice3 := make([]int, 2)
	slice4 := make([]int, 2, 10)
	fmt.Printf("%d，%d\n", len(slice3), cap(slice3))
	fmt.Printf("%d, %d\n", len(slice4), cap(slice4))

	var slice5 []int
	fmt.Println(slice5 == nil)
	// 插入一个元素
	slice5 = append(slice5, 0)
	fmt.Println(slice5)
	// 插入多个元素
	slice5 = append(slice5, 2, 3, 4)
	fmt.Println(slice5)
	// 插入一个数组，...是解包，不能省略
	slice5 = append(slice5, []int{5, 6, 7}...)
	fmt.Println(slice5)
	// 在第一个位置插入数组
	slice5 = append([]int{0}, slice5...)
	fmt.Println(slice5)
	// 在中间插入元素
	slice5 = append(slice5[:5], append([]int{10, 10}, slice5[5:]...)...)
	fmt.Println(slice5)
}

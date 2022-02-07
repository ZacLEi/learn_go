package main

import (
	"fmt"
)

/**
字典的使用
*/
func main() {
	var scores map[string]int = map[string]int{"english": 90, "chinese": 100}

	scores1 := map[string]int{"english": 90, "chinese": 100}

	scores2 := make(map[string]int)
	scores2["english"] = 90
	scores2["chinese"] = 100

	fmt.Println(scores)
	fmt.Println(scores1)
	fmt.Println(scores2)

	// 相关操作
	scores["math"] = 200 // 添加
	scores["math"] = 240 // 修改
	delete(scores, "math")
	// 如果要判断一个值是否存在
	gen, ok := scores["gen"]
	if !ok {
		fmt.Println("gen 不存在" + string(gen))
	}

	// 循环字典
	for key, value := range scores2 {
		fmt.Printf("%s 的值是：%d\n", key, value)
	}

	// 只循环key
	for key, _ := range scores2 {
		fmt.Printf("%s \n", key)
	}

	// 只循环value
	for _, value := range scores2 {
		fmt.Printf("%d\n", value)
	}
}

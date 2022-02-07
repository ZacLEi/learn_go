package main

import (
	"fmt"
	"strconv"
)

/**
接口实现多态：
只要实现了接口中的所有方法，我就认为这个结构体实现了这个接口
*/

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

// Phone
// 如果有一个类型/结构体，实现了一个接口要求的所有方法，
// 这里 Phone 接口只有 call方法，所以只要实现了 call 方法，我们就可以称它实现了 Phone 接口。
// 这个接口的实现是隐式的，不像 JAVA 中要用 implements 显示说明。
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

// 实现了 Good 接口要求的两个方法后，手机和赠品在Go语言看来就都是商品（Good）类型了。
// FreeGift
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func main() {
	// 这时候，我挑选了两件商品（实例化），分别是手机和耳机（赠品，不要钱）
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}
	// 然后创建一个购物车（也就是类型为 Good的切片），来存放这些商品。
	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}

// 定义一个方法来计算购物车里的订单金额
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

package main

import "fmt"

/**
结构体
*/

// Profile type Profile struct {
//	name   string
//	age    int
//	gender string
//	mother *Profile // 指针
//	father *Profile // 指针
//}
//若相邻的属性（字段）是相同类型，可以合并写在一起
type Profile struct {
	name, gender string
	age          int
	mother       *Profile // 指针
	father       *Profile // 指针
}

// FmtProfile 绑定方法 其中FmtProfile 是方法名，
//而(person Profile) ：表示将 FmtProfile 方法与 Profile 的实例绑定。
//我们把 Profile 称为方法的接收者，而 person 表示实例本身，
//它相当于 Python 中的 self，在方法内可以使用 person.属性名 的方法来访问实例属性。
func (person Profile) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

// 方法参数的传递方式
// 当你想要在方法内改变实例的属性的时候，必须使用指针做为方法的接收者。
// 重点在于这个星号: *
func (person *Profile) increaseAge() {
	person.age += 1
}

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company  // 匿名字段 通过组合实现继承
}

func main() {
	myself := Profile{name: "xiaoming", age: 24, gender: "mail"}
	fmt.Println(myself.age)
	myself.increaseAge()
	fmt.Println(myself.age)

	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company:  myCom,
	}

	fmt.Println(staffInfo.company.companyName)
	fmt.Println(staffInfo.companyName)

	// 内部方法和外部方法 首字母大写就是public方法 消协是private
	// 选择器： 当你的对象是个结构体的指针时，.省去*取值操作，直接获取结构体的属性值
	//p1:= &Profile{"123"}
	//fmt.Println(p1.name)
}

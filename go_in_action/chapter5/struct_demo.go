package chapter5

import "fmt"

// 自定义结构体一
type user struct {
	name       string
	age        int
	privileged bool
}

// 自定义结构体二（基于已有类型, 值互相兼容，不能互相赋值）
type Duration int64

func structDemo() {
	var shawock user
	shawock.name = "shawock"
	shawock.age = 23
	shawock.privileged = true
	fmt.Println(shawock)

	// var 创建一个变量并初始化为其零值
	var jenny user
	fmt.Println(jenny)

	// := 短变量声明操作符在一次操作中完成两件事情：声明一个变量，并初始化。
	// 用结构字面量来完成这个初始化方式一
	lisa := user{
		name:       "lisa",
		age:        31,
		privileged: true,
	}
	fmt.Println(lisa)

	// 用结构字面量来完成这个初始化方式二
	lisa2 := user{"lisa", 29, false}
	fmt.Println(lisa2)

	// 编译错误
	//var duration Duration
	//duration = int64(1000)
}

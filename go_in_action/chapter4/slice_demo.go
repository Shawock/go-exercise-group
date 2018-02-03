package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 5, 5)
	slice[0] = "hello"
	slice[1] = "slice"
	fmt.Println(slice)

	for k, v := range slice {
		fmt.Println(v == "")
		fmt.Printf("k = %d, v = %s\n", k, v)
	}

	fmt.Println()

	slice1 := make([]int, 5)
	slice1[0] = 99
	for k, v := range slice1 {
		fmt.Printf("k = %d, v = %d\n", k, v)
	}

	slice2 := []string{"red", "orange", "yellow", "green", "gray"}
	fmt.Println(slice2)

	slice3 := []int{99, 100, 105}
	fmt.Println(slice3)

	slice4 := []string{99: ""}
	fmt.Printf("len = %d, cap = %d\n", len(slice4), cap(slice4))

	var slice5 []int
	fmt.Printf("slice5 = %v\n", slice5 == nil)

	slice6 := []int{}
	fmt.Println(slice6)

	slice7 := make([]int, 0)
	fmt.Println(slice7)

	slice8 := []int{10, 20, 30, 40, 50, 60}
	// 从idx = 1 的位置开始切，切到第三个（不含）
	// len = 3 - 1, cap = len(slice8) - 1
	slice81 := slice8[1:3]
	fmt.Printf("slice8 = %v, slice81 = %v, len(slice81) = %v\n", slice8, slice81, cap(slice81))
	// 切片只能访问到其长度内的元素，那搞这个容易似乎没卵用
	//slice81[2] = 33 // panic: runtime error: index out of range
	// 与切片的容量相关联的元素只能用于增长切片。在使用这部分元素前，必须将其合并到切片的长度里。
	//fmt.Println(slice81[2])

	slice9 := []int{1, 2, 3, 4}
	slice91 := append(slice9, 50)
	fmt.Printf("slice9 = %v, slice91 = %v, len = %v, cap = %v\n", slice9, slice91, len(slice91), cap(slice91))

	slice10 := []int{1, 2}
	slice11 := []int{3, 4}
	slice1011 := append(slice10, slice11...)
	fmt.Printf("append result = %v\n", slice1011)

	// 现代 for 循环
	for index, value := range slice1011 {
		fmt.Printf("index = %d, value = %d, valueAddr = %X, eleAddr = %X\n", index, value, &value, &slice1011[index])
	}

	// 传统 for 循环
	for index := 0; index < len(slice1011); index++ {
		fmt.Printf("index = %d, value = %d\n", index, slice1011[index])
	}

	slice12 := [][]int{{10}, {20, 30}}
	fmt.Println(slice12)

	slice12[0] = append(slice12[0], 20)
	fmt.Println(slice12)

	slice13 := make([]int, 1e6)
	// 在函数间传递切片就是要在函数间以值的方式传递切片
	// 函数在调用和返回时都会复制切片
	slice1301 := foo(slice13)
	fmt.Println(&slice13 == &slice1301)
}

func foo(slice []int) []int {
	return slice
}

package main

import "fmt"

// 映射的键可以是任何值。这个值的类型可以是内置的类型，也可以是结构类型，只要这个值
// 可以使用==运算符做比较。切片、函数以及包含切片的结构类型这些类型由于具有引用语义，
// 不能作为映射的键

// 将切片或者映射传递给函数成本很小，并且不会复制底层的数据结构。
func main() {
	dict1 := make(map[string]int)
	fmt.Println(dict1)

	dict2 := make(map[string][]int)
	fmt.Println(dict2)

	dict3 := make(map[string][2]int)

	dict4 := map[int][]string{1: {"aaa"}, 2: {"bbb"}}

	fmt.Printf("%v, %v\n", dict3, dict4)

	colors := map[string]string{
		"red":    "aaaaa",
		"orange": "bbbbb",
		"yellow": "ccccc",
		"green":  "ddddd",
		"blue":   "eeeee",
	}

	printMap(colors)

	removeColor(colors, "blue")

	printMap(colors)

	fmt.Println(colors)

	fmt.Println(contains(colors, "black"))
	fmt.Println(contains(colors, "red"))
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

func printMap(dict map[string]string) {
	fmt.Printf("=============\ndict size = %d\n", len(dict))
	for k, v := range dict {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}
}

func contains(dict map[string]string, key string) bool {
	_, exist := dict[key]
	return exist
}

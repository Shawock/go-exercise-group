package main

import (
	"fmt"
	_ "fmt" // 下划线让编译器接受这类导入，并且调用对应包内的所有代码文件里定义的 init 函数
	"log"
	"os"
)

func init() {
	log.Println("init invoked.")
	log.SetOutput(os.Stderr)
}

func main() {
	log.Println("application start successfully.")
	log.Println(length("hello shawock"))
	fmt.Printf("The quick brown fox jumped over lazy dogs", 3.14)
}

func length(str string) int {
	return len(str)
}

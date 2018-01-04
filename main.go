package main

import (
	_ "fmt" // 下划线让编译器接受这类导入，并且调用对应包内的所有代码文件里定义的 init 函数
	"github.com/shawock/go-exercise-group/go-in-action/chapter5"
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

	counter := chapter5.AlertCounter(10)
	log.Println(counter)

	counter1 := chapter5.New(100)
	log.Printf("counter1's type = %T\n", counter1)

	admin := chapter5.AdminV2{Right: 3}
	log.Printf("%v\n", admin)

	admin.Name = "Shawock"
	admin.Email = "zhouhao@google.com"
	log.Printf("%+v\n", admin)
	log.Printf("%p\n", &admin)

	b := byte(1)
	log.Printf("%T\n", b)
}

func length(str string) int {
	return len(str)
}
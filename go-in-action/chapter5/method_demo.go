package chapter5

import (
	"fmt"
)

type person struct {
	name  string
	email string
}

// 如果一个函数有接收者，这个函数就被称为方法。
// Go语言既允许使用值，也允许使用指针来调用方法，不必严格符合接收者的类型。
// Go 编译器为了支持这种方法调用在背后做的事情，帮我们做了指针被解引用为值，或者 引用值得到一个指针
func (p person) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", p.name, p.email)
}

func (p *person) changeEmail(email string) {
	p.email = email
}

func (p person) changeEmailV2(email string) {
	p.email = email
}

// 下面编译不通过 编译器只允许为命名的用户定义的类型声明方法
//func (dict map[string]int) test(aaa string) {
//
//}

func methodDemo() {
	tom := person{"tom", "tom@google.com"}
	tom.notify()

	jerry := &person{"jerry", "jerry@google.com"}
	jerry.notify()

	tom.changeEmail("mot@google.com")
	tom.notify()

	jerry.changeEmail("yerry@google.com")
	jerry.notify()

	lucy := &person{"lucy", "lucy@google.com"}
	lucy.notify()
	lucy.changeEmailV2("lily@google.com")
	lucy.notify()
}

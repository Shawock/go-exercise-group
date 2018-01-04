package chapter5

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// demo1
type notifier interface {
	notify()
}

type Animal struct {
	name  string
	owner string
}

func (a *Animal) notify() {
	fmt.Printf("this animal %s belongs to %s", a.name, a.owner)
}

func sendNotification(n notifier) {
	n.notify()
}

// demo2
type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", d)
}

func interfaceDemo() {
	bs := []byte("hello")
	fmt.Println(len(bs))

	var b bytes.Buffer
	b.Write([]byte("hello "))
	fmt.Fprint(&b, "world")
	io.Copy(os.Stdout, &b)

	animal1 := Animal{"miumiu", "Alex"}
	// 下面这个是编译不通过的
	// sendNotification(animal1)
	sendNotification(&animal1)

	//下面这样是可以通过编译的? TODO
	//d := duration(123)
	//d.pretty()

	//下面这个无法运行时拿到指针? TODO
	//duration(123).pretty()
}

type Reader interface {
	Read(p []byte) (int, error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() (err error)
}
package pointer

import "fmt"

/**
	- 变量初始化后，该变量的指针（开辟的内存空间）位置不变，如果重新对此变量赋值，将刷新该内存空间中的内容
	- 无法对一个的值变量的指针做改变？ e.g. a := 1 / b := 2 / &a = &b 运行异常 ？？不太确定
	- & 取指针操作
	- * 解析指针操作
	- 类型前面加 * 表示持有该类型的指针
	- 如果有个结构体的指针, 如 demo3 中的 s 变量，要访问结构体的内容，不需要显示使用 * 解指针，直播调用，如 s.name
	- 方法形参是指针时，通过该指针可直接操作原有对象的内容，
	- 方法形参是值时，传入的是原值的 copy 值（指针发生变更）
 */
func main() {
	//demo0()
	//demo1()
	//demo2()
	//demo3()
}

func demo0() {
	a := 1
	fmt.Printf("%p -> %d\n", &a, a)
	a = 2
	fmt.Printf("%p -> %d\n", &a, a)

	b := 1
	//&a = &b 运行异常 ？
	fmt.Printf("a(%p -> %d), b(%p -> %d)\n", &a, a, &b, b)
}

func demo1() {
	i := 1
	var p *int
	p = &i
	fmt.Printf("i = %d; p = %v; *p = %d\n", i, p, *p)

	*p = 2
	fmt.Printf("i = %d; p = %v; *p = %d\n", i, p, *p)

	i = 3
	fmt.Printf("i = %d; p = %v; *p = %d\n", i, p, *p)
}

type abc struct {
	v int
}

func (a abc) method1() {
	fmt.Printf("method1 &a = %p\n", &a)
	a.v = 1
	fmt.Printf("1: %d\n", a.v)
}

func (a *abc) method2() {
	fmt.Printf("method2 &a = %p\n", a)
	fmt.Printf("2: %d\n", a.v)
	a.v = 2
	fmt.Printf("3: %d\n", a.v)
}

func (a *abc) method3() {
	fmt.Printf("method3 &a = %p\n", a)
	fmt.Printf("4: %d\n", a.v)
}

func demo2() {
	aObj := abc{}
	fmt.Printf("demo2 &a =  %p\n", &aObj)
	aObj.method1()
	aObj.method2()
	aObj.method3()
}

func demo3() {
	s := buildUser("Jack")
	fmt.Printf("s address = %v, s content = %v", &s, *s)
}

type S map[string][]string

func buildUser(name string) (s *S) {
	return &S{
		"name":           {name},
		"badges":         {"creator", "pgc"},
		"promoteRegions": {"CN", "US"},
	}
}

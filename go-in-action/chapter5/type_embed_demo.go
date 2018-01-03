package chapter5

import "fmt"

type userBasic struct {
	name string
}

func (u *userBasic) post() {
	fmt.Printf("%s posts new video\n", u.name)
}

// 注意声明字段和嵌入类型在语法上的不同
type liveUser struct {
	userBasic
	likeNum int
	name    string
}

func typeEmbedDemo() {
	lu := liveUser{userBasic{"baby"}, 123, "jason"}
	lu.userBasic.post()
	lu.post()

	fmt.Println(lu.userBasic.name)
	fmt.Println(lu.name)
}

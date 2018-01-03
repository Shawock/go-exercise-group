package chapter5

import "fmt"

type eraser interface {
	erase()
}

type musical struct {
	musicalId int
	title     string
}

func (m *musical) erase() {
	fmt.Printf("musical[id=%d] erased successfully, title = %s\n", m.musicalId, m.title)
}

type comment struct {
	commentId int
	text      string
}

func (c *comment) erase() {
	fmt.Printf("comment[id=%d] erased successfully, text = %s\n", c.commentId, c.text)
}

func eraseObject(e eraser) {
	e.erase()
}

func duotai() {
	musical1 := musical{11111, "hello muser"}
	// 关于这里为什么要用 & 不是太理解
	eraseObject(&musical1)

	comment1 := comment{22222, "you're so beautiful."}
	eraseObject(&comment1)
}

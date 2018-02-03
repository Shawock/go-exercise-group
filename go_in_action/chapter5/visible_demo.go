package chapter5

type alertCounter int

type AlertCounter int

//将工厂函数命名为 New 是 Go 语言的一个习惯。
func New(v int) alertCounter {
	return alertCounter(v)
}

type userV2 struct {
	Name  string
	Email string
}

type AdminV2 struct {
	userV2
	Right int
}

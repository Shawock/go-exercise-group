package search

// 空结构在创建实例时，不会分配任何内存。
type defaultMatcher struct {
}

type defaultMatcherV2 struct {
}

// 使用值作为接收者声明的方法，在接口类型的值为值或者指针时，都可以被调用。不懂？ TODO
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// 使用指针作为接收者声明的方法，只能在接口类型的值是一个指针的时候被调用。不懂？ TODO
func (m *defaultMatcherV2) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	panic("implement me")
}

func init() {
	var dm11 defaultMatcher
	dm11.Search(nil, "hello")

	dm12 := new(defaultMatcher)
	dm12.Search(nil, "hello")

	var dm21 defaultMatcherV2
	dm21.Search(nil, "hello")

	dm22 := new(defaultMatcherV2)
	dm22.Search(nil, "hello")

	var dm13 defaultMatcher
	var matcher Matcher = &dm13 // 将指针赋值给接口类型
	matcher.Search(nil, "hello")
}

package chapter7

import (
	"os"
	"time"
)

// runner 包用于展示如何使用通道来监视程序的执行时间，
// 如果程序运行时间太长，也可以用 runner 包来终止程序。

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
}
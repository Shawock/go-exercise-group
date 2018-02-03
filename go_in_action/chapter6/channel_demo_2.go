package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg2 sync.WaitGroup
)

func main() {
	wg2.Add(1)

	jilibang := make(chan int)

	go running(jilibang)

	jilibang <- 1

	wg2.Wait()
}

func running(jilibang chan int) {
	// 先接棒
	runner := <-jilibang
	fmt.Printf("runner get jilibang and start running %d\n", runner)

	// 接棒后开始跑
	time.Sleep(2 * time.Second)

	// 要不要传递棒子还是结束接力赛，根据当前是第几棒决定
	if runner == 5 {
		// 第 4 棒选手跑到终点结束
		fmt.Printf("runner finished %d\n", runner)
		wg2.Done()
	} else {
		// 将接力棒传递给下一棒选手
		newRunner := runner + 1

		// 下一棒选手做准备(注意：下一棒选手的 goroutine 虽然已经开始但是被锁住了，直到棒子传过来
		go running(jilibang)

		// 将接力棒传递给下一棒选手
		fmt.Printf("runner exchanged from %d to %d\n", runner, newRunner)
		jilibang <- newRunner
	}
}

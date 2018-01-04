package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	counter int
	swg     sync.WaitGroup
	// mutex 用来定义一段代码临界区
	mtx sync.Mutex
)

func main() {
	runtime.GOMAXPROCS(1)

	swg.Add(2)

	go incCounter(1)
	go incCounter(2)

	swg.Wait()
	fmt.Printf("final counter = %d", counter)
}

func incCounter(id int) {
	defer swg.Done()

	fmt.Printf("id = %d incCounter start.\n", id)
	for i := 0; i < 2; i++ {
		mtx.Lock()
		// 大括号不是必需的，只是让代码更清晰
		{
			value := counter
			// 互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界区代码
			// 当前 goroutine 从线程退出，并放回到队列,后面的 goroutine 就算执行到这里也没用，拿不到锁，无法执行
			runtime.Gosched()
			// 这个测试只是让他执行时间长点
			time.Sleep(1 * time.Second)
			value++
			counter = value
		}
		mtx.Unlock()
	}
}

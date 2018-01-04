package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutDown int32
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("aaa")
	go doWork("bbb")

	time.Sleep(1000 * time.Millisecond)

	atomic.StoreInt32(&shutDown, 1)
	wg.Wait()
	fmt.Println("main finish.")
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("doing work %s\n", name)

		if atomic.LoadInt32(&shutDown) == 1 {
			fmt.Printf("shutdown work %s\n", name)
			break
		}
	}
}

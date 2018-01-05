package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	tasks            = 10
)

var (
	wg3 sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	taskChan := make(chan string, tasks)

	wg3.Add(numberGoroutines)

	for i := 1; i <= numberGoroutines; i++ {
		go worker(taskChan, i)
	}

	for i := 1; i <= tasks; i++ {
		taskChan <- fmt.Sprintf("t[%d]", i)
	}

	close(taskChan)

	wg3.Wait()
}

func worker(taskChan chan string, id int) {
	defer wg3.Done()

	for {
		task, ok := <-taskChan

		if ok {
			// 处理任务
			fmt.Printf("worker: %d process task %s start\n", id, task)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("worker: %d process task %s done\n", id, task)
		} else {
			// 没有任务待处理
			fmt.Printf("worker: %d shutdown\n", id)
			break
		}
	}
}

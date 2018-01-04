package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg1 sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg1.Add(2)

	court := make(chan int)

	go player("Shawock", court)
	go player("Jenny", court)

	court <- 1

	//time.Sleep(5 * time.Second)

	wg1.Wait()

}

func player(name string, court chan int) {
	defer wg1.Done()

	for {
		// 这个接收动作会锁住 goroutine，直到有数据发送到通道里。
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won the game\n", name)
			return
		}

		r := rand.Intn(100)

		if r%17 == 0 {
			fmt.Printf("player %s miss the %dth ball\n", name, ball)
			close(court)
			return
		}

		fmt.Printf("player %s hit the %dth ball\n", name, ball)
		ball++
		court <- ball
	}
}

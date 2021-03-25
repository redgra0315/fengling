package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context
var wg sync.WaitGroup

//var notify bool
var exitChan chan bool = make(chan bool, 1)

func f(ctx context.Context) {
LOOP:
	for {
		fmt.Println("hello")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	defer wg.Done()
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 如何通知子grouttine退出
	//notify = true
	exitChan <- true
	cancel()
	wg.Wait()

}

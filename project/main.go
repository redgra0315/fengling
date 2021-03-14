package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// waitGroup

// math/rand
func f() {
	// 每次执行的时候输出的数字都不一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int(), rand.Intn(10))
	}
}
func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
	//defer wg.Done()
}

// GOMAXPROCS

func f3() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println("A", i)
	}
}
func f4() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println("b", i)
	}
}

var wg sync.WaitGroup

func main() {
	//f()
	//for i := 0; i < 10; i++ {
	//	wg.Add(i)
	//	go f1(i)
	//	//ch := make(chan int)
	//	//<-ch
	//}
	//wg.Wait()

	// 设置CPU的逻辑核心数，默认跑满整个CPU
	runtime.GOMAXPROCS(8)
	runtime.NumCPU()
	wg.Add(2)
	go f3()
	go f4()
	wg.Wait()

}

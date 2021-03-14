package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello world!", i)
	//time.Sleep(5 * time.Second)
}
func main() {
	// go 关键子开启一个单独的goroutine,
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		//go hello(i)
		fmt.Println("main")
		time.Sleep(100)
	}
}

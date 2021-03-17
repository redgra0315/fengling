package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
开启一个goroutine循环生成int64类型的随机数，发送到jobChan
开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
主goroutine从resultChan取出结果并打印到终端输出
*/
type job struct {
	value int64
}
type Result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func test1(te chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newjob := &job{
			value: x,
		}
		te <- newjob
		time.Sleep(time.Millisecond * 500)
	}
}
func test2(zl <-chan *job, resultChan chan<- *Result) {
	defer wg.Done()
	for {
		job := <-zl
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10

		}
		newResult := &Result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go test1(jobChan)
	for i := 0; i < 24; i++ {
		go test2(jobChan, resultChan)

	}
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()

}

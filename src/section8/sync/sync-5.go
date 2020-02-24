package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, notify, notifyALL : 기타 언어
	// Wait, Signal, Broadcast : Golang

	// ex1

	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting :", n)
			condition.Wait() // 고루틴 대기 (멈춤)
			fmt.Println("waiting End :", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		// fmt.Println("received : ", <-c)
	}

	// 고루틴 순서대로 깨움
	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Wake Goroutine(Signal) : ", i)
	// 	condition.Signal() // 한개씩 깨움 (모든 고루틴 생성 후)
	// 	mutex.Unlock()
	// }

	// 고루틴 한번에 깨움
	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}

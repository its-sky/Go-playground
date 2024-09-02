package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// Go는 기본적으로 1개의 CPU를 사용하는데 복수개의 CPU를 사용할 경우 아래와 같이 호출하여야 한다
	runtime.GOMAXPROCS(4)

	// 동기적 실행
	say("Sync")

	// 비동기적 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	time.Sleep(time.Second * 3)

	// WaitGroup 생성. 2개의 고루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 이용한 goroutine
	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 끝나면 .Done()호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait()
}

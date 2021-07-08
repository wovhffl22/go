package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++ // 1초 간격으로 0부터 시작 1씩 증가하는 인티져 (정수형)
	}
}

func main() {
	c := make(chan int) // channel 생성

	go push(c) // Channel을 push함수에 c인자로 넘기고 go Thread시작

	timerChan := time.After(10 * time.Second) // 특정시간 주기적 반환하는 타이머 채널을 만든다

	for {
		select {
		case v := <-c: // c에 입력값이 있으면 아래 v출력
			fmt.Println(v)
		case <-timerChan: // 만든타임채널을 주기적으로 살펴본다
			// case <-time.After(10 * time.Second):	// 선언문 새로운 타이머 채널 생성을 반복
			fmt.Println("timed out")
			return
		}
	}
}

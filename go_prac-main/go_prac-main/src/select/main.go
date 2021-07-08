package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(2 * time.Second)
		c <- i
		i++ // 2초 간격으로 0부터 시작 1씩 증가하는 인티져 (정수형)
	}
}

func main() {
	c := make(chan int) // Channel을 push함수에 c인자로 넘기고 go Thread시작

	go push(c)

	for {
		select {
		case v := <-c: // c에 입력값이 있으면 아래 v출력
			fmt.Println(v)
		default: // c에 입력값이 없거나 일반적일 때에는 아래 Idle출력
			fmt.Println("Idle")
			time.Sleep(1 * time.Second) // 1초 쉬면서 Idle출력
		}
	}
}

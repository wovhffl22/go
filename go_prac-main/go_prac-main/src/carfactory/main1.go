package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string // Plane 만들기
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) { //channel
	for { //무한루프
		select { //타이어공장
		case car := <-carChan:
			car.val += "Tier_C. " //카형 타이어
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, " //비행기형 타이어
			outPlaneChan <- plane   //출력함수
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for { //무한루프
		select { //
		case car := <-carChan:
			car.val += "Engine_C. " //카형 엔진
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, " //비행기형 엔진
			outPlaneChan <- plane     //출력함수
		}
	}
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i)} // int를 string으로 변환함수
		i++
	}
}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(carChan1, planeChan1, carChan2, planeChan2)
	go MakeEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
	}

}

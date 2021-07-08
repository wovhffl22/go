package main

import "fmt"

func main() {
	//method1()
	//receiver()

	PVMethod()
}

// <method>
type account struct {
	balance int
}

func withdrawFunc(p *account, amount int) {
	p.balance -= amount // p가 가지고 있는 주소의 값을 감소
}

func (p *account) withdrawMethod(amount int) { // 내부에서 withdrawFunc()함수로
	p.balance -= amount //
}

func method1() {
	a := &account{100}   // 주소를 받아 저장 p
	withdrawFunc(a, 30)  // a에 담긴 p주소의 값을 감소
	a.withdrawMethod(30) // 내부적으로 위의 함수로 대체
	fmt.Printf("%d \n", a.balance)
}

// <receiver>
// 기본 데이터형들도 메서드를 가진것 처럼 만들 수 있다
type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func receiver() {
	var b myInt = 10
	fmt.Println(b.add(31))
}

// PointerMethod ValueTypeMethod
type account1 struct {
	balance   int
	firstName string
	lastname  string
}

// 포인터 메서드
func (a1 *account1) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account1) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account1) withdrawReturnValue(amount int) account1 {
	a3.balance -= amount // mainA의 복사본
	return a3            // return
}

func PVMethod() {
	var mainA *account1 = &account1{100, "Bruce", "Lee"} // mainA(포인터)

	mainA.withdrawPointer(30)  // (a1 *account1, 30)
	fmt.Println(mainA.balance) //

	mainA.withdrawValue(20)    //	a2 연산 후 GC
	fmt.Println(mainA.balance) //

	var mainB account1 = mainA.withdrawReturnValue(20) // mainB(구조체)
	fmt.Println(mainB.balance)                         //	return a3을 mainB에 저장

	mainB.withdrawPointer(30) // (*mainB).(30)
	fmt.Println(mainB.balance)
}

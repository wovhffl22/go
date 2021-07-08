package main

import "fmt"

//	> go mod init typecasting
//	> go build
//	> go run cast

func main() {

	first()

	second()
}

func first() {
	a := 3               // 64 정수
	var b float32 = 3.14 // 부호 있는 실수

	var c int = int(b)  // 강제 형변환, 연산자 int()
	d := float32(a) * b // 의도하지 않은 형 변환을 예방하여 런타임 에러를 예방 // float32() 수정
	// 원하는 타입으로 타입 캐스팅하여 연산을 하면 된다
	var e int16 = 16
	f := a + int(e)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
}

func second() {
	var a int16 = 3456
	var c int8 = int8(a)

	fmt.Println(a) // 3456	// 2진수, 00001101 10000000 // 상위바이트, 하위바이트 // little Endian
	fmt.Println(c) // -128	// 양의 a가 타입캐스팅 문제로 값이 변할 수 있다.
}

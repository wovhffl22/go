package main

import (
	"fmt"
)

//  > go mod init ascii
//  > go build
//  > go run ascii

func main() { // 메인 패키지는 반드시
	// var ch1 byte = 65
	// var ch2 byte = 0101 //
	// var ch3 byte = 0x41 //
	// var ch4 byte = 'A'  //

	// fmt.Printf("\n %d %c %c %c \n", ch1, ch2, ch3, ch4) // 4개의 변환? 4개의 변수

	// var uni1 rune = 44032
	// var uni2 rune = 0126000
	// var uni3 rune = 0xAC00
	// var uni4 rune = '가'

	// fmt.Printf("\n %c %c %c %c \n", uni1, uni2, uni3, uni4)
	// PrintFormat

	// rawLiteral := `아리랑\n
	// 	아리랑\n
	// 		아라리오`

	// interLiteral := "\n아리랑아리랑\n아라리오"

	// // fmt.Print(rawLiteral)
	// // fmt.Print(interLiteral)

	// fmt.Println()
	// fmt.Print(unsafe.Sizeof(rawLiteral)) // 16바이트로 출력
	// fmt.Println()
	// fmt.Print(unsafe.Sizeof(interLiteral)) // 문자열로 저장 할 때에는 16바이트 사이즈의 주소를 저장
	// fmt.Println()
	// cpu 연산 과정
	// string 결국 문자열의 문자 주소를 저장하여
	// 메모리 힙 공간의 10kbyte의 열자리 주소 공간을 찾아 main주소의 erp? 전기적 신호로
	// 컴파일 한 시점에 상수화 되어 주소에 박힌다
	// print 함수에 for문을 포함하고 있어
	// for if str == '(10자리)' %c로 출력하고 주소를++ 증가하여.. \0을 만날 때 까지 // \0 이스케이프 문자
	// 메모리에는 12자리의 문자열 주소만 저장이 됨
	// 데이터를 원하는 형태로 핸들링 하기 위해 - 포인터

	// var a int
	// var b int

	// // 함수의 결과값을 변수에 대입하라
	// n, err := fmt.Scan(&a, &b) // &a &주소
	// if err != nil {
	// 	fmt.Println(n, err) //
	// } else {
	// 	fmt.Println(n, a, b)
	// }
	// 버퍼 공간에 값이 누적되기 때문에 재 실행

	// Scan (c *int, d *int)
	// c = 주소 d = 주소
	// *c = 값  *c의 가리키는 주소에 값을 저장?
	//
	// Scanf의 버퍼 스캔영역은 OS레벨이기에
	// 입력값이 파이프에 적제된다고 생각 stringpipe

	// import bufio os

	// studio := bufio.NewReader(os.Stdin)

	// var a int
	// var b int

	// n, err := fmt.Scanln(&a, &b)

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("다시 입력해 주세요")
	// 	studio.ReadString('\n') // 넘겨진 인자값 까지 버퍼를 비워 버린다
	// } else {
	// 	fmt.Println(n, a, b)
	// }

	// n, err = fmt.Scanln(&a, &b) // := 새로운 변수가 아니어서 에러발생

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("다시 입력해 주세요")
	// 	studio.ReadString('\n')
	// } else {
	// 	fmt.Println(n, a, b)
	// }
	// 버퍼의 잘못 된 값을 비워서 에러를 방지하는 중요성

	// var i = 10
	// var j = 20
	// fmt.Println(i, j)
	// swap(&i, &j) //

	// fmt.Println(i, j)
	// //
	// const pi = 3.141592 // 상수

	// const (
	// 	red   int = iota // iota = 0 / 예약어
	// 	blue  int = iota // iota = 1
	// 	green            // green = iota // = 2
	// )
	// //
	// d := "Hello, " + "world!" // "_ __" // "_의 주소를 리턴 \n까지 반복

	// a := 112                // 01110000 / 56 28 14
	// b := a >> 3             // 00001110: a의 비트를 왼쪽으로 3번 이동
	// fmt.Printf("%08b\n", b) // 00001110
	// fmt.Printf("%d\n", b)

	//

	const (
		Red float32 = float32(iota)
		Green
	)

	fmt.Printf("%d, %d", Red, Green)

	// var j float64 = 0.1
	// var h float64 = 0.2

	// fmt.Println(1 == 1)             // true: 두 정수가 같으므로 true
	// fmt.Println(3.5 == 3.5)         // true: 두 실수가 같으므로 true
	// fmt.Println("Hello" == "Hello") // true: 두 문자열이 같으므로 true
	// fmt.Println(j+h == 0.3)
	// fmt.Println(j + h)

	// a := [3]int{1, 2, 3}
	// b := [3]int{1, 2, 3}
	// fmt.Println(a[1])
	// fmt.Println(a == b) // true: 두 배열이 같으므로 true

	// c := []int{1, 2, 3}
	// fmt.Println(c == nil) // false: 슬라이스를 nil과 비교하여
	// // 메모리가 할당되었는지 확인

	// d := map[string]int{"Hello": 1}
	// fmt.Println(d == nil) // false: 맵을 nil과 비교하여
	// // 메모리가 할당되었는지 확인

	// e := 1
	// var p1 *int = &e
	// var p2 *int = &e
	// fmt.Println(p1 == p2) // true: 포인터에 저장된 메모리 주소가 같으므로 true

}

// call by pointer	// reference(x)
// func swap(a *int, b *int) { // &(x) 컴파일 오류
// 	*b, *a = *a, *b
// 	// int temp = *a
// 	// *a = *b
// 	// *b = temp 기존의 c 방식
// 	//
// }

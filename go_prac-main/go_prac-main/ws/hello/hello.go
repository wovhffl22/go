package main

import (
	"fmt" // double quotantion
	"unsafe"
)

// 터미널 실행, 경로 이동
// > dir			디렉터리
// > cd ws/hello	hello디렉터리 이동

//	go.mod 파일 생성
//	> go mod init hello		/	현재 디렉터리를 모듈의 루트로 설정
//	> go build				/	Compile하여 실행 파일 생성
//	> go run hello.go		/	.go 파일 실행
//	> ./hello

func main() { // 다음 줄 { 시작 시 Error
	fmt.Println("Hello World")

	var a int = 3
	var c = 4
	d := 2 // 선언 대행문, 정수 선언문으로 인식

	var aa int8 = 1
	var bb int16 = 1
	var cc int32 = 1
	var dd int64 = 1

	//	변수 크기 구하기
	fmt.Println(unsafe.Sizeof(aa)) // unsafe패키지의 Sizeof 함수를 사용
	fmt.Println(unsafe.Sizeof(bb))
	fmt.Println(unsafe.Sizeof(cc))
	fmt.Println(unsafe.Sizeof(dd))
	fmt.Println()

	// 원래 변수들
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))

	fmt.Println(a, c, d)
}

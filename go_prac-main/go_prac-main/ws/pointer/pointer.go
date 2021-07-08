package main

import "fmt"

type Data struct { // data 충돌 없음, 예방 차원
	value int
	data  [100]int
}

//*int와 같이 주소는 동일한 필드의 길이를 가진다, 5byte(40bit)?
func ChangeData(arg *Data) { // (변수 *타입)
	arg.data[99] = 5 //	구조체가 있기 때문에 * & 없이 Data필드를 찾음
	arg.value = 3    // (*arg).data[99] = 0 구조체가 아닐 경우
}

func main() {
	//first()
	//second()
	//third()

	fourth()
}

func fourth() {
	str4 := "AAAA"
	str5 := "aaBBB"
	//str6 := "BBB"
	//str7 := "cccccccccc"
	str10 := str4 + str5
	fmt.Println(str10)

	// 문자 단위로 비교, 맨 앞의 문자가 큰 값이면 이후의 비교는 pass
	fmt.Printf("%s > %s : %v\n", str4, str5, str4 > str5)
}

func third() {
	//
	var ch rune = '한'
	var ech rune = 'A'
	// rune	내부적으로 int32로 저장
	// type rune = int32
	// 저장되는 정수를 몇 바이트로 볼 것인가 - type
	fmt.Println(ch, ech)            // 정수
	fmt.Printf("%c, %c\n", ch, ech) //	문자
	fmt.Printf("%T, %T\n", ch, ech) //	타입

	//
	str1 := "안녕하세요"                                                      // 3byte
	str2 := "ABCDE"                                                      //	1byte
	fmt.Printf("len(str1) = %d, len(str2) = %d\n", len(str1), len(str2)) // 길이(str1)

	//
	str3 := "hello 월드"
	for n, val := range str3 { // index,value	// array, slice, map 가능
		fmt.Printf("index : %d, value : %c\n", n, val)
	}

}

func second() {
	var data Data
	ChangeData(&data) // & 앰퍼샌드 자연스럽게 사용 중요
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[99] = %d\n", data.data[99])
}

func first() {
	var numPtr *int
	fmt.Println(numPtr)
}

// > go mod init main
// > go build
// > go run pointer.go
// <nil>	// 정의되지 않은 메모리의 영역을 표시
//
// var i int
// var p *int = &i	// i변수 주소를 포인터 타입 p변수에 넣음

package main

import "fmt"

// C, JAVA의 형태
// void IncreaseAndReturn() {
// }

//	> go mod init condition
//	> go build
//	> go run condition

func main() {

	conditional()

	uploadFile()

	localVariable()

	switch_()
}

// conditional
func conditional() {

	// conditional
	// temp := 25
	// peer := 80

	// if peer, temp := rainPercent80; peer >= 25 {
	// 	fmt.Println("덥고 습합니다.")
	// } else {
	// 	fmt.Println("야외 활동하기 좋습니다.")
	// }

	tmp := 30
	if tmp >= 28 {
		fmt.Println("에어컨을 켜세요.")
	} else if tmp < 5 {
		fmt.Println("에어컨을 끄세요.")
	} else {
		fmt.Println("에어컨이 동작중입니다.")
	}
	fmt.Println("종료!!")
}

// func rainOver80() {
// 	//강수확률 80% 이상
// 	fmt.Println("덥고 비가 옵니다.")
// }
// func rainOver20() {
// 	fmt.Println("덥고 습합니다.")
// }
// func rainUnder20() {
// 	fmt.Println("야외 활동하기 좋습니다.")
// }
// func rainPercent80(p int) bool {
// 	if p > 80 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// conditional2
func UploadFile() (filename string, success bool) {
	filename = "hello.exe" // 문자열은 리터럴로 대체 된다
	success = true
	return // 리턴을 자동 완성 모든 타입을 리턴
}

func uploadFile() {

	//if초기문 초기화문, 함수호출 을 동시에
	if filename, success := UploadFile(); success { // 판단식 만으로 success의 타입을 유추
		// if문 실행과 동시에 결과를 리턴
		fmt.Println("Upload success", filename)
	} else {
		fmt.Println("Upload Failed", filename)
	}
}

// localVariable
// 전역변수	--> Data Segment
var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func localVariable() {
	// 지역변수	-->	Stack Segment
	var i int = 10
	i = IncreaseAndReturn()               // 특정 코드 내에서 바로 호출 가능
	if false && IncreaseAndReturn() < 5 { // 판단문
		fmt.Println("1 증가", i)
	}
	if true || IncreaseAndReturn() < 5 { // 판단문
		// (true || IncreaseAndReturn() < 5) //(x)
		fmt.Println("1 증가", i)
	}

}

// Recursive function
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)          // 재귀함수 자기함수 호출
	fmt.Println("After", n) // 321After123 ?
}

func recursive_() {
	printNo(3)
}

// Recursive function
func factorial(i int) int { // 선언이 명시적이면 리턴을 생략 가능, 아닐 경우 리턴 명시
	// (j int, k int) 복수일 경우 () 필수
	// a,b := func factorial.. 앞의 변수가 bool일 경우 조건문으로
	// 가변인수?
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func recursive() {

	//	재귀함수
	var f int
	fmt.Println("계산하고자 하는 팩토리얼 값을 입력하세요")
	fmt.Scanf("%d", &f) // 중요 앰퍼샌드 사용

	fmt.Printf("%d 팩토리얼 값은 ... %d", f, factorial(f))
}

// switch
func switch_() {
	//
	var a int = 95
	var b int = 88

	switch true {
	case b > 90, (a + 1) > 90: // || ,는 or 역할
		fmt.Println("모두 다 A학점입니다")
	case a > 80, b > 80:
		fmt.Println("모두 다 B학점입니다")
	default:
		fmt.Println("모두 다 F학점입니다")

	}
}

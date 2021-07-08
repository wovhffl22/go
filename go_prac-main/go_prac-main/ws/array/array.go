package main

import (
	"fmt"
	"sort"
)

// 배열 초기화
func main() {
	//	배열의 원소 개별[첨자]은 변수명으로 인식.
	array()

	//arrfor()

	//arrprintf()

	//day()

	//scanprint()

	//buffer()

	//forRange()

	//forLable()

	//arrSlice()

	//arrSort()

	//arrCopy()

	//map_()

}

func map_() {
	// golang의 디폴트 자료구조
	var m map[int]string
	m = make(map[int]string)
	// 추가 혹은 갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	// 키에 대한 값 읽기
	str := m[134]
	fmt.Println(str)

	noData := m[999] // 값이 없으면 nill 혹은 zero 리턴
	fmt.Println(noData)

	var n map[string]int
	n = make(map[string]int)

	n["익명"] = 10
	value, exists := n["익명"] // "익명" 키에 값을 넣음?
	fmt.Println(value, exists)

	value, exists = n["범인"]      // 초기 선언시에만 :
	fmt.Println(n["범인"], exists) // Default

	delete(n, "익명") // 초기화 ~리턴x
	value, exists = n["익명"]
	fmt.Println(value, exists)
}

func arrCopy() {
	s := []int{10, 3, 21, 7, 41, 1}
	t := make([]int, 4, 20)

	copy(t, s) // range와 다르게 사이즈가 작으면 작은 만큼 만 복사가 이루어짐
	fmt.Println(t)
}

func arrSort() {
	s := []int{10, 3, 21, 7, 41, 1}

	sort.Sort(sort.IntSlice(s)) // 오름차순
	fmt.Printf("%d\n", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s))) //	내림차순
	fmt.Printf("%d", s)

}

func arrSlice() {

	s := []int{10, 11, 21, 31, 41, 51}
	t := make([]int, 7, 7) // len(s) //s //s[2:5]
	fmt.Printf("%v, %v\n", s, t)
	fmt.Printf("%p, %p\n", s[0:1], t[0:1]) // 슬라이스 하여 주소를 출력, 첫 슬라이스의 주소?

	for n, val := range s { // s슬라이스를 반복하여, n의 위치로
		t[n] = val
	}
	fmt.Printf("%v, %v\n", s, t) // t의 길이 만큼

	var b [3]int
	fmt.Printf("%x\n", b) // 원소의 주소

	var a []int // 배열
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s = make([]int, 5, 10) // 타입, 길이, 용량
	fmt.Println(s, len(s), cap(s))
}

func forLable() {
	i := 0

L1:
	for {
		for {
			if i == 0 {
				break L1 //	소속된 반복문을 벗어남
			}
		}
		fmt.Println("Inner OK")
	}
	fmt.Println("Outer Label OK")
	fmt.Println("OK")
}

func forRange() {
	var names []string = []string{"이순신", "홍길동", "강감찬"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for n, val := range names { // _, val / n, _ // 명시적 묵시 // 인덱스, 값
		fmt.Println(n, val) // 함수 배열 호출과 유사
		// (n, val) / (n) / (val)
	}
}

func buffer() {
	var i int = 10
	fmt.Scanf("%d", &i)
	fmt.Printf("%v\n", &i) // 인자에 맞는 타입으로 변환하여 줌
	fmt.Println(i)

}

func scanprint() {

	var day [7]string

	for i := 0; i < len(day); i++ {
		fmt.Scanf("%s", &day[i]) //
	}
	for i := 0; i < len(day); i++ {
		fmt.Printf("%s요일", day[i])
	}
}

func day() {
	var day = [7]string{"월", "화", "수", "목", "금", "토", "일"}

	for j := 0; j < len(day); j++ {
		fmt.Printf("%s요일, ", day[j])
		//	day[첨자]의 타입 string
		//	원자의 첫(월) 주소를 찾아가 값을 가져옴.. 반복하여 마지막 원자의 \0을 만나 종료
	}
}

func arrprintf() {
	//
	var a = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Printf("%.2f, %.2f, %.2f, %.2f, %.2f\n", a[0], a[1], a[2], a[3], a[4])
}

func arrfor() {
	//반복문 형태
	// for i := 0; i < 4; i++ {
	// 	fmt.Printf("%.2f, ", a[i])
	// }
	// fmt.Printf("%.2f", a[4])
}

func array() {

	var a = [2][3]int{ //	타입의 개념으로 생각
		{1, 2, 3},
		{4, 5, 6}, // 콤마 추가
	}

	var b = [3]int{
		1, 2, 3,
	}

	var c = [2][2][3]int{
		{ //
			{1, 2, 3},
			{4, 5, 6},
		},

		{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	var d = [4][2][2][3]int{
		{ //	[4]
			{ //	[2]
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			{ //
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			{ //
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			{ //
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

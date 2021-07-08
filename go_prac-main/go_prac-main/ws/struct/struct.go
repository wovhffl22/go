package main

import "fmt"

type person struct {
	name string // , 없이
	age  int
}

type User struct { // 첫 대문자로 외부로 노출여부 설정
	Name string // 현재 패키지를 벗어나 외부에 노출 된다
	age  int    // 외부에 알려지지 않음
	Id   string //
}

type VIPUser struct {
	UserInfo User // public 역할
	VIPLevel int
	Price    int
}

// 메인 함수의 외부
//	함수의 외부에 할당문은 syntax error
// p4 := person{}	// 할당x
// var p4 person	// 전역변수 선언은 가능

func main() {

	//structSecond()
	structFirst()
}

func structSecond() {
	normalUser := User{Name: "이순신", age: 55, Id: "Lee"}                  // 내부에서는 노출여부 접근 가능
	vipUser := VIPUser{User{Name: "이순신", age: 55, Id: "Lee"}, 5, 100000} // 구조체로 대입 가능
	fmt.Println(normalUser, vipUser)
	fmt.Printf("유저 이름 : %s, 나이 : %d, 아이디 : %s\n", normalUser.Name, normalUser.age, normalUser.Id) // 싱글 구조체 접근
	fmt.Printf("VIP유저 이름 : %s, ", vipUser.UserInfo.Id)

	vipUser2 := vipUser // go언어는 주소 복사가 아니라 같은 값을 복사
	fmt.Println(vipUser2)
}

func structFirst() {
	var p1 person
	p2 := person{age: 30, name: "Lee"} // 필드에 값을 모두 넣어야 함
	//p2 := person{}
	// var p1 person
	// var p2 person

	p3 := new(person) // 힙 공간에 주소를 할당
	// var p3 *person

	fmt.Println(p1.age, p2.age, p3.age)
	fmt.Printf("%p, %p, %p\n", &p1, &p2, p3)
	// 메모리에 억세스할 경우 포인터로 동작

}

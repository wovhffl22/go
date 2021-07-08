package main

import "fmt"

type Data struct {
	book string
	rank int
}

func ChangeData(arg *Data) {
	arg.kook = 90
	arg.young = 92
	arg.soo = 82
}

func max() {
	mySlice := make([]Data, 0)
	mySlice = append(mySlice, Data{"국어", 90})
	mySlice = append(mySlice, Data{"영어", 92})
	mySlice = append(mySlice, Data{"수학", 82})

	fmt
}

func main() {

	var data Data

	// i := [2][2][3]int{
	// 	{
	// 		{0, 1, 2},
	// 		{3, 4, 5},
	// 	},
	// 	{
	// 		{0, 1, 2},
	// 		{3, 4, 5},
	// 	},
	// }

	// a := [5]byte{'A', 'B', 'C', 'D', 'E'}
	// for _, v := range a {
	// 	fmt.Printf("%s", string(v))
	// }
}

func weather() {
	var temp, per int

	fmt.Scanf("%d %d", &temp, &per)

	if tempOver25(temp) { //	if peer, temp := rainPercent80; peer >= 25 {
		if rainPercent80(per) {
			rainOver80()
		} else if rainPercent20(per) {
			rainOver20()
		} else if !rainPercent20(per) {
			rainUnder20()
		}
	} else {
		if temp < 10 || rainPercent80(per) {
			fmt.Println("야외 활동하기 좋지 않습니다.")
		}
		fmt.Println("좋은 날씨입니다.")
	}
	// if i, j := string bool; bool {}
	// if peer, temp := rainPercent80; peer >= 25 {
	// 	//fmt.Println("덥고 습합니다.")
	// } else {
	// 	fmt.Println("야외 활동하기 좋습니다.")
	// }
}

func rainOver80() {
	//강수확률 80% 이상
	fmt.Println("덥고 비가 옵니다.")
}
func rainOver20() {
	fmt.Println("덥고 습합니다.")
}
func rainUnder20() {
	fmt.Println("야외 활동하기 좋습니다.")
}

func rainPercent80(p int) bool {
	if p >= 80 {
		return true
	} else {
		return false
	}
}

func rainPercent20(p int) bool {
	if p >= 20 {
		return true
	} else {
		return false
	}
}

func tempOver25(t int) bool {
	if t >= 25 {
		return true
	} else {
		return false
	}
}

// 퀴즈
// 일기예보에서 낮 최고 기온이 25도 이상이고 강수확률이 80% 이상이면 "덥고 비가 옵니다." 출력,
// 낮 최고 기온이 25도 이상이고 강수확률이 20% 이상이면 "덥고 습합니다."
// 낮 최고 기온이 25도 이상이고 강수확률이 20% 미만이면 "야외 활동하기 좋습니다" 출력
// 기온이 25도 이상이 아니고 기온이 10도 미만이거나 강수확률이 80% 이상이면 "야외 활동하기 좋지 않습니다" 그렇지 않다면 "좋은 날씨입니다" 출력

// 단, 강수확률 80% 이상 함수의 이름은 rainOver80..
// 강수확률 20% 이상 함수의 이름은 rainOver20..
// 강수확률 20% 미만 함수의 이름은 rainUnder20.. 함수를 만들어서 if문의 초기화문에 대입하여 코딩하세요

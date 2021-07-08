package main

import "fmt"

type Bread struct {
	val string
}
type OrangeJam struct {
	opened bool
}
type SpoonOfOrange struct {
}
type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)

	for i := 0; i < num; i++ {

		breads[i] = &Bread{val: "bread"}

	}

	return breads
}
func OpenOrangeJam(jam *OrangeJam) {

	jam.opened = true
}
func GetOneSpoon(_ *OrangeJam) *SpoonOfOrange {

	return &SpoonOfOrange{}
}
func PutJamOnBread(bread *Bread, jam *SpoonOfOrange) {

	for i := 0; i < 2; i++ {

		bread.val += " + orange Jam"

	}
}
func MakeSandwitch(breads []*Bread) *Sandwitch {

	sandwitch := &Sandwitch{}

	for i := 0; i < len(breads); i++ {

		sandwitch.val += breads[i].val + " + "
	}

	return sandwitch

}

func main() {
	breads := GetBreads(5)

	jam := &OrangeJam{}

	OpenOrangeJam(jam)

	spoon := GetOneSpoon(jam)

	for i := 0; i < len(breads); i++ {

		PutJamOnBread(breads[i], spoon)

	}

	sandwitch := MakeSandwitch(breads)

	fmt.Println(sandwitch.val)

}

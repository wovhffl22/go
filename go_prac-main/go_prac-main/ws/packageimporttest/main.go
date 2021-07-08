package main

// GOPATH GOPATH 둘 경로에서 찾지 못 함
import "kerygma.com/greeting"

// kerygma.com/greeting => ../greeting
// > go get kerygma.com/greeting
func main() {
	greeting.SayHello()
}

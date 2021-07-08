package main

// import "..." ??
import "kerygma.com/structPackage"

func main() {
	// ... ??
	original_sp := structPackage.StructPackage{}
	structPackage.ShowPackageInfo(&original_sp, "stopBlock", "v1.0.1") // &주소의 필드를 "","" 로
	// 동적 struct 구조체
}

// > go mod init main
// > go get kerygma.com/structPackage
// > go build
// > go run main.go

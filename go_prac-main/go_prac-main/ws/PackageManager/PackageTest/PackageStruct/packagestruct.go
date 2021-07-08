package structPackage

import "fmt"

type StructPackage struct {
	Name    string
	Varsion string
}

func ShowPackageInfo(sp *StructPackage, name string, ver string) { // sp *StructPackage, "stopBlock", "v1.0.0"
	//sp := StructPackage{Name: "goBlock", Varsion: "v1.0.0"}
	//sp = new(StructPackage)
	sp.Name = name // heap에 할당하지 않아도
	sp.Varsion = ver
	fmt.Printf("Package Name : %s, Version : %s\n", sp.Name, sp.Varsion)
}

// go mod init structPackage

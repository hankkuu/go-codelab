package main

import (
	"fmt"
)

// import "fmt" // 이런 형태도 가능

func main() {
	fmt.Println("hello world")
}

// terminal: go run ModuleName/FolderName or cd helloworld -> go run .
// 파일 저장하고 구동해야 함
// 사용하지 않는 package가 import 되면 컴파일이 되지 않음

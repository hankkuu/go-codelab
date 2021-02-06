package main

import (
	"fmt"
	"reflect"
)

// 전역변수
var (
	// 아래의 경우는 안됨
	//g := 3333
	g        = 3333
	UserName = "Kim" // export 대상에는 대문자로 시작
)

// 전역상수
const (
	UserAge = 20
)

func main() {
	var a int
	fmt.Println(a)

	// (컴파일시) 타입추론을 위해 초기화와 동시에 할당
	var b = 0
	fmt.Println(reflect.TypeOf(b))

	var c = "11111"
	fmt.Println(reflect.TypeOf(c))

	var d = int32(11111)
	fmt.Println(reflect.TypeOf(d))

	e := int32(2222)
	fmt.Println(reflect.TypeOf(e))

	//var f
	//fmt.Println(reflect.TypeOf(f))

	const h int = 10 // := 상수는 이거 안됨
	//h = 100

	var (
		i int    = 10
		j string = "22"
	)
	fmt.Println(i)
	fmt.Println(j)

	var k, l int32 = 10, 20
	fmt.Println(k + l)

	// 기본적으로 카멜케이스를 사용
	userName := "kim geon"
	fmt.Println(userName)
}

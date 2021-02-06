package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func main() {
	var intType int = 10
	fmt.Printf("%v %T\n", intType, intType)

	var float32Type float32 = 3.14
	fmt.Printf("%v %T\n", float32Type, float32Type)

	var stringType string = "Hello world"
	fmt.Printf("%v %T\n", stringType, stringType)

	fmt.Println(intType + int(float32Type))

	var intpType *int
	fmt.Printf("%v %T\n", intpType, intpType) // <nil> *int // null 상태

	intpType = &intType // 포인트 참조의 형태로 사용 가능 단 포인트 연산은 사용할 일이 그닥 없다 (unsafe.Point)
	fmt.Println(&intType)
	fmt.Println(intpType)

	*intpType = 100
	fmt.Println(intType)

	type user struct {
		name string
	}

	user1 := User{
		Name: "aa",
	}

	fmt.Println(reflect.TypeOf(user1))
}

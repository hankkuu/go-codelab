package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := Any("123FF")
	fmt.Println(a)

	a = 123
	fmt.Println(a)

	PrintV(1)
	//PrintV(true)
	//PrintV([]int{1, 2, 3})
	PrintSwitch("string") // 내부에서 type 별로 switch 하는 방식으로 권장 사항
}

// 어떠한 매서드도 필요하지 않은 인터페이스 (즉 어떤 값도 다 들어갈 수 있다 : 만능변수)
type Any interface {
}

func Print() {
	fmt.Println()
}

func PrintV(v interface{}) {
	// 타입이 모르니 되지 않음
	//v += 1

	// 이렇게 했을 때 위의 경우는 나오지 않지만 v가 string 이면 오류
	a := v.(int)
	fmt.Println(reflect.TypeOf(a))

	a, ok := v.(int)
	if !ok {
		fmt.Println(a)
		a = 100
	}

	fmt.Println(v)
}

func PrintSwitch(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println(val + 1)
	case string:
		fmt.Println(val + " 123")
	default:
		fmt.Println("invalid type")
	}
}

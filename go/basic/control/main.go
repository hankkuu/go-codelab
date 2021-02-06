package main

import (
	"bytes"
	"fmt"
)

// if, switch, for
func main() {

	var ok bool = true
	if ok {
		fmt.Println("OK")
	}

	var num int = 0
	//if num {
	if num != 0 {
		fmt.Println("Not")
	}

	// str은 if 안에서만 사용
	if str := "Hello World"; str == "Hello World" {
		fmt.Println(str)
	} else {
		fmt.Println("else: ", str)
	}

	//str = "change"
	// err은 안에서만 사용
	if n, err := bytes.NewBuffer(nil).Read(nil); err != nil {
		fmt.Println(n)
	}

	n, err := bytes.NewBuffer(nil).Read(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// 기본적으로 break 가 사용되어 있어 아래로 흐르지 않음 fallthrough로 해야 함
	str := "hello"
	switch str {
	case "hello":
		fmt.Println(123)
		fallthrough
	case "world":
		fmt.Println(456)
	default:
		fmt.Println("default")
	}

	str2 := "hello"
	switch {
	case str2 == "hello":
		fmt.Println(123)
		fallthrough
	case str2 == "world":
		fmt.Println(456)
	default:
		fmt.Println("default")
	}

	ok = false
	for ok {

	}

	// 무한루프
	//for {
	//}

	// i는 for 문에서만 사용
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 조건만 걸어 for문을 사용할 수도 있다
	for str == "ok" {

	}

}

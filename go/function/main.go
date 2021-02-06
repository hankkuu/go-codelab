package main

import (
	"fmt"
)

func main() {
	function1(1, 2)
	fmt.Println("func2: ", function2(1, 2))
	val1, val2 := function3(1, 2)
	fmt.Println("func3: ", val1, val2)

	// 이 방식은 되지만 지양하는게 좋을 듯
	val1, val2 = function4(1, 2)
	fmt.Println("func4: ", val1, val2)

	// 함수도 변수처럼 사용할 수 있다
	function4 := func(a int, b int) (int, int) {
		return a * b, a / b
	}
	fmt.Println(function4(1, 2))

	// callback 함수를 만들 수 있다
	function6(function4)
}

func function1(a int, b int) {
	fmt.Println("func1: ", a, b)
}

func function2(a int, b int) int {
	return a + b
}

func function3(a int, b int) (int, int) {
	return a + b, a - b
}

func function4(a int, b int) (c int, d int) {
	c = 100
	d = 1000
	return
}

func function5(a int, b int) (int, int) {
	var c, d int
	return c, d
}

func function6(f func(int, int) (int, int)) {
	fmt.Println(f(1, 2))
}

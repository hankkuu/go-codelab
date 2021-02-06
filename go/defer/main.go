package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	function1()
	// defer 함수는 함수가 끝날때 stack에 쌓아놓고 꺼내 쓰는 것처럼 역순으로 호출이 된다 (지연된 호출)
	defer function1()

	function2()
	defer function2()

	function3()
	defer function3()
}

// 파일 입출력에서 close를 해줄 수 있다 go 는 try catch finally 가 없기 때문에 finally 부분에서 리소스 해제를 해줄 수 없다
func read() {
	f, err := os.Open("../go.mod")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

}

func function1() {
	fmt.Println("1")
}

func function2() {
	fmt.Println("2")
}

func function3() {
	fmt.Println("3")
}

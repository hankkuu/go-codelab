package main

import (
	"fmt"
	"log"
)

// er 을 붙이는 네이밍이 선호된다
type Stringer interface {
	String() string
	Error() string
}

type Name string

// implements 역할을 하는데 따로 상속관계가 없다
func (n Name) String() string {
	return string(n)
}

func main() {
	name := Name("Geon")
	PrintString(name)

	//error
	LogError(name + " 1 ")
}

func (n Name) Error() string {
	return string(n)
}

func LogError(err error) {
	log.Println(err)
}

func PrintString(s Stringer) {
	fmt.Println("stringer: ", s.String())
	fmt.Println("stringer: ", s.Error())

}

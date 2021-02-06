package main

import (
	"fmt"
)

type User struct {
	ID   string
	Name string
}

func main() {
	user := User{
		ID:   "id1234",
		Name: "Geon Kim",
	}
	fmt.Printf("original : %p\n", &user)
	user.Print()
	user.ChangeID("new Id")
	user.Print()

	User.Print(user)
	(*User).ChangeID(&user, "id1")

	// struct 타입이 아닐 때 매서드 사용
	// 단, 외부 패키지에 정의된 내용은 가져와서 매서드로 만들지는 못한다
	name := Name("test")
	fmt.Println(name)
	fmt.Printf("value method : %p\n", &name)
	name.Print()

}

// value method - value가 넘어가면서 새로 복사한 개체를 사용하게 된다
// 매서드 리시버 이름 : u
func (u User) Print() {
	fmt.Printf("value method : %p\n", &u)
	fmt.Println("ID: ", u.ID)
	fmt.Println("Name: ", u.Name)
}

// point method - 주소가 넘어가면서 기존의 개체를 사용하게 된다
func (u *User) ChangeID(newId string) {
	fmt.Printf("point method : %p\n", u)
	//(*u).ID = newId
	u.ID = newId
}

type Name string

func (n Name) Print() {
	fmt.Printf("value method : %p\n", &n)
	fmt.Println("Name: ", n)
}

package main

import (
	"codelab/basic/struct/sub"
	"fmt"
)

type User struct {
	Name string
	Age  int
	Info Info
}

type Info struct {
	Phone string
}

type Admin struct {
	User
	Permistion int
}

func main() {

	user := sub.User{}
	fmt.Println(user)

	// int default = 0
	user1 := sub.User{
		Name: "Geon",
	}
	fmt.Println(user1)

	user2 := sub.User{
		Name: "Geon",
		Age:  20,
		//phone: "010-1111-1111", // 소문자 속성은 외부 접근이 안된다
	}
	fmt.Println(user2)
	fmt.Println(user2.Age)

	user3 := User{
		Name: "Geon",
		Age:  20,
		Info: Info{
			Phone: "010-1111-1111",
		},
	}
	fmt.Println(user3)
	fmt.Println(user3.Info.Phone)

	admin1 := Admin{
		User: User{
			Name: "Geon",
			Age:  20,
			Info: Info{
				Phone: "010-1111-1111",
			},
		},
	}
	//json.Marshal(admin1)

	fmt.Println(admin1)
	fmt.Println(admin1.Name) // User로 접근안하고 사용이 된다
	fmt.Println(admin1.Info)

}

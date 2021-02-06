package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

// Array, Slice, Map
func main() {

	// 할당하지 않으면 기본값 설정
	var a [3]int = [3]int{9}
	fmt.Println(a[0], a[1])

	//a.append(a, 3)

	fmt.Println("!!!!!!!!!!!", reflect.TypeOf(a))

	b := [2]int{1, 2}
	fmt.Println(b)

	c := [...]int{4, 6}
	fmt.Println(c[0], c[1], len(c))

	// for 문
	for idx, val := range c {
		fmt.Println(idx, val)
	}

	// string의 경우는 기본값이 "" empty로 된다
	d := [3]string{"hello", "!!!"}
	fmt.Println(d, len(d))

	///////////////////////////////////
	// slice - arr[시작인덱스(포함):끝인덱스(불포함)] 으로 슬라이싱 가능
	// e: Array -> Slice
	e := d[1:]
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))

	// 가변길이로 만들 수 있다 (배열은 고정된 길이를 이용한다면 가변길이로 만들면 슬라이싱할 대상을 조절할 수 있다)
	f := make([]int, 0)
	fmt.Println(f)

	f = append(f, 10)
	f = append(f, 20)
	f = append(f, 30)
	f = append(f, 40)
	f = append(f, 50)
	fmt.Println(f, len(f))

	g := f[:3]
	fmt.Println(g)

	fmt.Println("!!!!!!!!!!!", reflect.TypeOf(g))

	var h []int
	for i := 0; i < 10; i++ {
		h = append(h, i)
	}
	fmt.Println(h, len(h), cap(h))
	// cap: capacity
	// Slice 내부구조: h = [Ptr][Len: size][Cap: size]

	i := []int{1, 200, 3000, 4000}
	fmt.Println(i, len(i), cap(i))

	// slice 으로 크기가 변경될 수 있다
	j := i[1:]
	fmt.Println(j, len(j), cap(j))

	// 같은 배열을 공유하지만 크기는 다르게 slicing 했다
	k := i[:3]
	fmt.Println(k, len(k), cap(k))
	k[0] = 100 // slice 의 대상인 배열을 공유하기 때문에 변화를 만들어 버렸다 (버그생성 가능성 높음)
	fmt.Println(i, len(i), cap(i))

	// slice 를 만들때는 make를 써라 (array를 통해 만들면 원본을 readonly로만 써야 할듯)
	l := make([]int, 0, 3) // 길이가 0, 용량이 3인 slice 생성
	fmt.Println(l, len(l), cap(l))

	l = append(l, 3)
	l = append(l, 2)
	l = append(l, 67)
	l = append(l, 7) // capa 가 동적으로 늘어난다 (여유공간을 미리 만드는 것으로 보인다)
	fmt.Println(l, len(l), cap(l))

	var slice []int           // slice는 기본값이 nil
	fmt.Println(slice == nil) // true

	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)

	for idx, val := range slice1 {
		fmt.Println(idx, val)
	}

	/////////////////
	/// Map
	// map도 slice와 같이 선언만 하면 nil 로 초기화 된다 (array가 0이나 empty 문자열로 기본값이 생성되는것과 차이 점 )
	var map1 map[string]int // key : string, value: int
	fmt.Println(map1 == nil)

	map2 := make(map[string]int)
	fmt.Println(map2, len(map2))
	map2["key1"] = 10
	map2["key2"] = 20
	fmt.Println(map2, len(map2)) // 2
	fmt.Println(map2["key2"])    // 20
	fmt.Println(map2["key3"])    // 0

	val, ok := map2["key3"]
	fmt.Println(val, ok) // 0 false

	map3 := map[string]int{
		"key1": 100,
		"key2": 200,
	}
	fmt.Println(map3)

	map4 := map[User]int{
		User{}: 100,
	}
	fmt.Println(map4)

	// 이거 순서가 항상 정열된게 아니다 random 하게 출력된다 (순서가 상관 없을 때 사용해야 할 듯)
	for key, val := range map3 {
		fmt.Println(key, val)
	}
}

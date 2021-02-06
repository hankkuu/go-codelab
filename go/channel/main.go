package main

import (
	"fmt"
	"time"
)

// Channel은 Goroutine 간의 통신이 필요할 때
func main() {
	// 생성
	// c := make(chan int)

	// go func() {
	// 	c <- 10 // 입력, 채널에 값을 넣는다, 채널에 입력을 하려면 goroutine에서 해야 한다
	// 	c <- 20
	// 	c <- 30
	// }()

	// fmt.Println(<-c) // 출력 // 값이 없으면 block 되고 위의 go func 에서 입력 되면 출력된다
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// 위에는 버퍼없이 값을 입출력 했고 아래는 버퍼 사용 (버퍼사이즈 만큼 넣어놓기 때문에 값 출력을 기다리는 block을 방지할 수 있다)
	// 관리할 자료형이 추가되기 때문에 block 걱정이 되는 경우에 사용할 경우 좋고 아니면 관리 포인트가 늘어난다
	bufferdC := make(chan int, 3)
	go func() {
		bufferdC <- 40

		bufferdC <- 50
		bufferdC <- 60

		//fmt.Println(<-bufferdC)

		//bufferdC <- 70
		fmt.Println("aaa")

		//bufferdC <- 80
	}()

	//bufferdC <- 30 // buffer size 초과 (deadlock!!)
	time.Sleep(2 * time.Second)
	// 명시적으로 닫아주려면 close
	close(bufferdC)

	//fmt.Println(<-bufferdC)
	//fmt.Println(<-bufferdC)
	//fmt.Println(<-bufferdC)

	// channel도 range 사용이 가능 단 range 안에서 channel이 닫힐 때까지 deadlock이 된다
	for val := range bufferdC {
		fmt.Println(val)
	}

	//close(bufferdC)
}

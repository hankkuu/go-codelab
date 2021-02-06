package main

import (
	"fmt"
	"strconv"
)

// 여러 채널에서 들어오는 값을 효과적으로 제어 하기 위함
func main() {
	taskChanA := make(chan int)
	taskChanB := make(chan string)
	taskChanC := make(chan bool)

	go taskA(taskChanA, 10)
	go taskB(taskChanB, 10)
	go taskC(taskChanC, 10)

	//label1:
	for {
		select {
		case val := <-taskChanA:
			fmt.Println("A: ", val)
		case val := <-taskChanB:
			fmt.Println("B: ", val)
		case val := <-taskChanC:
			fmt.Println("C: ", val)
		// default 가 없으면 deadlock이 걸린다 어떠한 채널에서도 값을 받지 못하는 경우를 처리한다
		// 하지만 goroutine은 순서를 보장 받지 못하고 바로 실행 될 수 있기 떄문에 잠시 대기와 같은 처리르 해줘야 한다
		default:
			//break label1 // 임시로
			//fmt.Println("Default")
			//time.Sleep(1 * time.Second)
			//return		// 프로그램이 출력이 다 되기 전에 조기 종료
			break // 무한루프
		}
	}

}

func taskA(c chan int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}
}

func taskB(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- strconv.Itoa(i)
	}
}

func taskC(c chan bool, n int) {
	for i := 0; i < n; i++ {
		c <- i%2 == 0
	}
}

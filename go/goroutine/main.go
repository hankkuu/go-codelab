package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	// go의 경량 쓰레드 go의 런타임에 의해 관리된다(커널 쓰레드 상위에 있다) // 10만개의 고루틴을 생성
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go Print(&wg)
	}

	// go routine이 실행되기 전에 main이 종료 되어 출력이 안됨
	// go routine 간 반복되는 함수호출간 block이 되지 않기 때문에 (동시성 호출이 됨) 출력을 보고 싶으면 sleep이 메인에 필요
	//time.Sleep(2 * time.Second)

	// done 이 호출이 다 될 때까지 기다리게 된다 즉 위에는 2초간 기다리고 끝나는 거라면 아래는 10개의 done이 모두 호출되고 wait 가 끝나게 된다
	wg.Wait()
}

func Print(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	// res, err := http.Get("https://example.com")
	// if err != nil {
	// 	panic(err)
	// }
	//  fmt.Println(res.StatusCode)

	fmt.Println("hello")
}

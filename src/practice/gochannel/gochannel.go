package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	println(i)

	/**
	Go 채널은 수신자와 송신자가 서로를 기다리는 속성을 지니고 있다.
	이를 이용하여 Goroutine이 끝날 때까지 기다리는 기능을 구현할 수 있다
	어떤 Go 루틴 작업이 실행되고 있을 때 메인 루틴은 <-done에서 계속 수신하며 대기한다.
	익명함수 goroutine에서 작업이 끝난 후, done 채널에 true를 보내면 수신자는 이를 받고 종료하게 된다.
	*/
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	<-done

	/**
	채널의 종류 : Unbuffered Channel, Buffered Channel
	위에서 봤던 예제는 Unbuffered의 종류로 하나의 수신자가 데이터를 받을 때까지
	송신자가 데이터를 보내는 채널에 묶여 있게 된다.

	Buffered Channel을 이용하면 수신자가 받을 준비가 되어 있지 않을 지라도
	지정된 버퍼만큼 데이터를 보내고 다른 일을 할 수 있다.
	버퍼 채널은 make(chan type, N)을 통해 생성하고 N에 버퍼 개수를 지정한다.
	*/

	/**
	버퍼 사이즈를 지정해주지 않으면 main 루틴에서 채널에 1을 보냈는데
	이를 수신할 goroutine이 없기 때문에 에러를 발생시킨다.
	버퍼 사이즈를 지정하여 버퍼에 데이터를 보낼 수 있게 되어 에러가 발생하지 않는다.
	*/
	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)

	ch1 := make(chan string, 1)
	sendChannel(ch1)
	close(ch1)

	// 채널을 닫아도 버퍼에 있는 데이터를 수신할 수 있다
	receiveChannel(ch1)

	/**
	Channel Range
	채널에서 송신자가 송신을 한 후 채널을 닫을 수 있다.
	수신자는 임의의 갯수의 데이터를 채널이 닫힐 때까지 수신할 수 있다.
	수신자는 채널이 닫히는 것을 체크하면서 계속 루프를 돈다.
	*/
	ch2 := make(chan int, 2)

	ch2 <- 1
	ch2 <- 2

	close(ch2)

	// 방법 1 : 채널이 닫힌 걸 감지할 때까지 계속 수신
	/*for {
		if i, success := <-ch; success {
			println(i)
		} else {
			break
		}
	}*/

	// 방법 2 : 위와 동일한 기능인데 더 간결한 표현
	for i := range ch2 {
		println(i)
	}

	/**
	Channel Select
	복수의 채널들을 기다리며 준비된 채널을 실행하는 기능을 제공한다.
	여러 개의 case문에서 각각 다른 채널을 기다리다가 준비 된 채널 case를 실행하는 것이다.
	복수 채널에 신호가 오면 Go 런타임이 랜덤하게 한 개를 선택한다.
	default 문이 있으면 기다리지 않고 이를 실행한다.

	*/

	/**
	여기서 break EXIT은 EXIT문으로 가서 EXIT을 다시 실행하는게 아니라
	기존 실행하던 컨텍스트의 다음 문장부터 실행함.
	*/
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)
EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

// 송신 채널은 chan<- 으로 지정할 수 있음
func sendChannel(ch chan<- string) {
	ch <- "Data"
}

// 수신 채널은 <-chan 으로 지정할 수 있음
func receiveChannel(ch <-chan string) {
	data := <-ch
	fmt.Println(data)

	if _, success := <-ch; !success {
		println("더이상 데이터 없음.")
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

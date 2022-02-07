package main

import (
	"fmt"
	"time"
)

/**
了解一下select case
*/
func main() {
	// 仅能用于信道的相关操作
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeout := make(chan bool, 1)
	c2 <- "hello"
	go makeTimeout(timeout, 5)

	select {
	case msg1 := <-c1:
		fmt.Println("c1 接收到了消息：", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 接收到了消息：", msg2)
	case timeout <- true:
		fmt.Println("timeout . exit.")
		//default:
		// 避免造成死锁，当select遍历完所有的case没有命中时，就会进入default分支，但是没写default的话，就会阻塞
		// 避免这种情况，可以设置一个超时时间
	}

	c3 := make(chan int, 2)
	c3 <- 2
	select {
	case c3 <- 2:
		fmt.Println("c3 received:", <-c3)
		fmt.Println("c3 received:", <-c3)
	default:

	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

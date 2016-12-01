package main

import (
	"fmt"
	"time"
)

func testgo() {
	go sayhelloMan() //通过go来调用协程执行方法
	go sayhelloWomen()

	time.Sleep(time.Millisecond) //由于goroutine是协程，所以这边如果不加主线程等待时间，main函数会直接结束，不会等待goroutine执行完毕

	fmt.Println("结束")
}

func testchannel() {
	ch := make(chan int) //定义一个channle，用来接收传递的值

	go send(ch, 10) //如果注释此行或者下面一行中的其中一行，都将会出现死锁报错 fatal error: all goroutines are asleep - deadlock!

	recive(ch)
}

func testbufferchannel() {
	//ch := make(chan int) //定义一个channle，用来接收传递的值
	bufferchannel := make(chan int, 1) //定义缓冲为1的channel
	bufferchannel <- 1                 //可以直接往缓冲的channel中存值，只要小于1
	//bufferchannel <- 2 //大于1，会报死锁的错误
	//ch <- 1 //对于无缓冲的channle直接赋值，会报死锁的错误

	fmt.Println(<-bufferchannel)
}

func testrangeandclose() {
	bufferchannel_1 := make(chan int, 10)

	go setNumber(cap(bufferchannel_1), bufferchannel_1)
	for i := range bufferchannel_1 { //通过range取缓冲的channel中的值
		fmt.Println(i)
	}
}

func testselect() {
	ch := make(chan int)   //定义一个channle，用来接收传递的值
	done := make(chan int) //用来传递完毕的标识

	go speaknumber(ch, done)

	for {
		select {
		case i := <-ch:
			fmt.Println("speack:", i)
		case <-done:
			fmt.Println("speack:done")
			return
		}
	}
}

func testovertime() {
	ch := make(chan int) //定义一个channle，用来接收传递的值

	go saynumber(ch)

	for {
		select {
		case i := <-ch:
			fmt.Println("say:", i)
		case <-time.After(1 * time.Second):
			fmt.Println("say:timeout")
			return
		}
	}
}

func saynumber(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func main() {

	testgo()

	testchannel()

	testbufferchannel()

	testrangeandclose()

	testselect()

	testovertime()
}

func send(c chan int, i int) {
	fmt.Println("发送消息：", i)
	c <- i //将值塞到channle中，如果recive方法没有执行，send方法所在协程将阻塞
}

func recive(c chan int) {
	i := <-c
	fmt.Println("接收消息：", i)
}

func setNumber(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c) //通过close关闭channel
}

func sayhelloMan() {
	for i := 0; i < 5; i++ {
		fmt.Println("你好，大兄弟")
		time.Sleep(10)
	}
}

func sayhelloWomen() {
	for i := 0; i < 5; i++ {
		fmt.Println("你好，大妹子")
		time.Sleep(10)
	}
}

func speaknumber(ch chan int, done chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- 0
}

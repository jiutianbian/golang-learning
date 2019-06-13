package main

import (
	"fmt"
	"time"
)

//课程：并发

//--------------------------示例一：简单的协程示例------------------------------//

//goroutine是Go并行设计的核心。goroutine说到底其实就是协程，它比线程更小，
//十几个goroutine可能体现在底层就是五六个线程，
//Go语言内部帮你实现了这些goroutine之间的内存共享。
//执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
//也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。
//goroutine是通过Go的runtime管理的一个线程管理器。
//goroutine通过go关键字实现了，请看如下示例1

func testgo() {
	go sayhelloMan()
	go sayhelloWomen()

	//由于goroutine是协程
	//所以这边如果不加主线程等待时间，main函数会直接结束，不会等待goroutine执行完毕
	time.Sleep(time.Millisecond)

	fmt.Println("结束")
}

func sayhelloMan() {
	for i := 0; i < 100; i++ {
		fmt.Println("你好，大兄弟")
		time.Sleep(10)
	}
}

func sayhelloWomen() {
	for i := 0; i < 100; i++ {
		fmt.Println("你好，大妹子")
		time.Sleep(10)
	}
}

//--------------------------示例二：协程通过channle通信示例------------------------------//
//goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。
//那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。
//channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。
//这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。
//注意，必须使用make 创建channel，请看如下示例2

func testchannel() {
	//1.我们先定义一个channle，用来接收传递的值
	ch := make(chan int)

	go send(ch, 10)

	//recive(ch)

}

func send(c chan int, i int) {
	fmt.Println("发送消息：", i)
	//channel通过操作符<-来接收和发送数据
	c <- i //将值塞到channle中，如果recive方法没有执行，send方法所在协程将阻塞
}

func recive(c chan int) {
	//channel通过操作符<-来接收和发送数据
	i := <-c
	fmt.Println("接收消息：", i)
}

//--------------------------示例三：协程通过缓冲channle通信示例------------------------------//
//上面我们介绍了默认的非缓存类型的channel，不过Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素
//ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。
//在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
//请看如下示例3

func testbufferchannel() {
	bufferchannel := make(chan int, 1) //定义缓冲为1的channel
	bufferchannel <- 1                 //可以直接往缓冲的channel中存值，只要小于1
	//bufferchannel <- 2 //大于1，会报死锁的错误

	fmt.Println(<-bufferchannel)
}

//--------------------------示例四：协程通过缓冲channle通信示例------------------------------//
//对于缓冲channle，我们可以可以通过range，像操作slice或者map一样操作缓存类型的channel，请看下面的示例4
func testrangeandclose() {
	bufferchannel_1 := make(chan int, 10)

	go setNumber(cap(bufferchannel_1), bufferchannel_1)
	for i := range bufferchannel_1 { //通过range取缓冲的channel中的值
		fmt.Println(i)
	}
}

//记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
//另外记住一点的就是channel不像文件之类的，不需要经常去关闭
//只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
func setNumber(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c) //通过close关闭channel
}

//我们上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢，
//Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行
//当多个channel都准备好的时候，select是随机的选择一个执行的。

//--------------------------示例五：select ------------------------------//
func testselect() {
	ch := make(chan int)   //定义一个channle，用来接收传递的值
	done := make(chan int) //定义一个channle，用来传递完毕的标识

	go speaknumber(ch, done)

	//相当于while循环
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

func speaknumber(ch chan int, done chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- 0
}

//func testovertime() {
//	ch := make(chan int) //定义一个channle，用来接收传递的值

//	go saynumber(ch)

//	for {
//		select {
//		case i := <-ch:
//			fmt.Println("say:", i)
//		case <-time.After(1 * time.Second):
//			fmt.Println("say:timeout")
//			return
//		}
//	}
//}

//func saynumber(ch chan int) {
//	for i := 0; i < 10; i++ {
//		ch <- i
//	}
//}

func main() {

	testgo()

	// testchannel()

	// testbufferchannel()

	//testrangeandclose()

	// testselect()

	//	testovertime()
}

package main

import (
	"fmt"
)

//课程：interface

//--------------------------示例一：interface简单示例------------------------------//

//各个编程语言关于接口的定义：接口相当于是一份契约，它规定了一个对象所能提供的一组操作。
//golang中的interface也是此作用：interface是就是一组method签名的组合，
//我们可以通过interface来定义对象的一组行为。
//在golang中，假如一个对象实现了interface中的所有method
//那么我们就说此对象实现了此interface
//通过interface我们可以让面向对象和内容组织实现非常的方便。

//1.我们先声明一个叫motion的接口
type motion interface {
	run()
	stop()
}

//2.我们再申明三个 对象 car 和 bike，hunman
type car struct {
	brand string
	price string
	speed string
	size  string
}

type bike struct {
	brand string
	price string
	speed string
}

type human struct {
	name string
	age  string
}

//3.car 和 bike 都实现了run和stop方法，我们就说car和bike实现了motion这个interface
//4.而 human 没有实现run方法，human就没有实现motion这个interface
func (c car) run() {
	fmt.Println("我是汽车:", c.brand, c.price, c.speed, "我跑起来了")
}

func (c car) stop() {
	fmt.Println("我是汽车:", c.brand, c.price, c.speed, "我停下来了")
}

func (b bike) run() {
	fmt.Println("我是自行车:", b.brand, b.price, b.speed, "我跑起来了")
}

func (b bike) stop() {
	fmt.Println("我是自行车:", b.brand, b.price, b.speed, "我停下来了")
}

func (h human) stop() {
	fmt.Println("我是:", h.name, h.age, "我停下来了")
}

func main() {
	//定义一个类型为motion的interface变量fc
	var fc motion

	var benchi car
	var fenghaung bike
	var yaoming human

	benchi = car{"奔驰", "60万", "140km/h", "很大"}
	fenghaung = bike{"凤凰", "1k", "20km/h"}
	yaoming = human{"姚明", "40"}

	//上面说过car和bike已经实现了motion这个interface方法
	//human没有实现motion这个interface方法
	//所以fc能够持有benchi和fenghaung，但不能持有yaoming
	fc = benchi
	fc.run()
	fc.stop()

	fc = fenghaung
	fc.run()
	fc.stop()

	//这个持有了会报错
	fc = yaoming
	yaoming.stop()

	//	testNullInterface()

	//	testCommaInterface()

	//	testParamInterface()

}

//--------------------------示例二：空interface简单示例------------------------------//
//空interface(interface{})不包含任何的method
//正因为如此，所有的类型都实现了空interface
//空interface功能有点类似java里object类

type nullInterface interface{}

func testNullInterface() {
	var parent nullInterface

	var age int
	var name string
	var yaoming human

	name = "易建联"
	age = 15
	yaoming = human{"姚明", "40"}

	//jiaobaba这个空的interface变量能够持有任意类型
	parent = name
	fmt.Println(parent)

	parent = age
	fmt.Println(parent)

	parent = yaoming
	fmt.Println(parent)

}

//--------------------------示例三：interface中类型判断示例------------------------------//
//我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。
//那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？通过Comma-ok断言
//Go语言里面有一个语法，可以直接判断是否是该类型的变量：
//value, ok = element.(T)，这里value就是变量的值，
//ok是一个bool类型，element是interface变量，T是断言的类型。
//示例如下
func showInterfaceType(nl nullInterface) {
	value, ok := nl.(int)

	fmt.Println("value:", value, "ok:", ok)
	//如果是调用showInterfaceType("1")，打印value: 0 ok: false
	//如果是调用showInterfaceType(1)，打印value: 1 ok: true
}

func testCommaInterface() {
	showInterfaceType("1")
	showInterfaceType(1)
}

// motion这个interface作为参数，可以接受实现了他的对象car和bike的变量
//func showYourFunction(fc motion) {
//	fc.run()
//	fc.stop()
//}

//func testParamInterface() {
//	benchi := car{"奔驰", "60万", "140km/h", "很大"}
//	fenghaung := bike{"凤凰", "1k", "20km/h"}

//	showYourFunction(benchi)
//	showYourFunction(fenghaung)
//}

package main

import (
	"fmt"
)

//声明一个叫function的接口
type function interface {
	run()
	stop()
}

//申明三个对象 car 和 bike，hunman
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

//car 和 bike 都实现了run和stop方法，我们就说car和bike实现了function这个interface，而 human 没有实现run方法，就没有实现function这个interface
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

type nullInterface interface{}

func testNullInterface() {
	var jiaobaba nullInterface

	var age int
	var name string
	var yaoming human

	name = "易建联"
	age = 15
	yaoming = human{"姚明", "40"}

	//jiaobaba这个空的interface变量能够持有任意类型
	jiaobaba = name
	fmt.Println(jiaobaba)

	jiaobaba = age
	fmt.Println(jiaobaba)

	jiaobaba = yaoming
	fmt.Println(jiaobaba)

}

// function这个interface作为参数，可以接受实现了他的对象car和bike的变量
func showYourFunction(fc function) {
	fc.run()
	fc.stop()
}

func testParamInterface() {
	benchi := car{"奔驰", "60万", "140km/h", "很大"}
	fenghaung := bike{"凤凰", "1k", "20km/h"}

	showYourFunction(benchi)
	showYourFunction(fenghaung)
}

func showInterfaceType(nl nullInterface) {
	value, ok := nl.(int)

	fmt.Println("value:", value, "ok:", ok)
	//打印value: 0 ok: false  value: 1 ok: true
}

func testCommaInterface() {
	showInterfaceType("1")
	showInterfaceType(1)
}

func main() {
	//定义一个类型为function的变量fc
	var fc function

	var benchi car
	var fenghaung bike
	//	var yaoming human

	benchi = car{"奔驰", "60万", "140km/h", "很大"}
	fenghaung = bike{"凤凰", "1k", "20km/h"}
	//	yaoming = human{"姚明", "40"}

	//fc能够持有实现了function这个interface方法的对象car和bike的变量，benchi和fenghaung，不能持有human
	fc = benchi
	fc.run()
	fc.stop()

	fc = fenghaung
	fc.run()
	fc.stop()

	//	fc = yaoming
	//	yaoming.stop()

	testNullInterface()

	testParamInterface()

	testCommaInterface()

}

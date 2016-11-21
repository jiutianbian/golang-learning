package main

import (
	"fmt"
)

//用struct来定义一个实体car
type car struct {
	brand string
	price string
	speed string
}

type suvcar struct {
	car  // struct作为匿名字段
	size string
}

//定义关联car的函数，这个函数为car这个实体的method，面向对象所需要的基本属性和方法定义这样就都有了
func (c car) show() {
	fmt.Println("我是:", c.brand, c.price, c.speed)
}

//重写 suvcar自己的show方法
func (c suvcar) show() {
	fmt.Println("我是:", c.brand, c.price, c.speed, c.size)
}

//指针作为接受者,比起直接传对象，可以用来保证设置属性的操作等真正作用于对象本身，而不是对象的copy
func (c *car) setBrand(brand string) {
	c.brand = brand
}

func (c *car) setPrice(price string) {
	c.price = price
}

func (c car) setSpeed(speed string) {
	c.speed = speed
}

func main() {
	var baoma car
	baoma.brand = "宝马"
	baoma.price = "40万"
	baoma.speed = "120km/h"
	baoma.show()

	benchi := new(car)
	benchi.setBrand("奔驰")
	benchi.setPrice("60万")
	benchi.setSpeed("140km/h")
	benchi.show()
	//输出 “我是: 奔驰 60万” 没有输出速度，因为setSpeed不是作用在car本生，而是作用在car的copy上

	var suvbaoma suvcar
	suvbaoma = suvcar{car{"宝马suv", "80万", "160km/h"}, "很高"}

	// 实现了匿名继承的struct，也能继承他的method
	suvbaoma.show()
	// 如果没有重写show方法，输出  我是: 宝马suv 80万 160km/h
	// 如果重写了show方法，输出  我是: 宝马suv 80万 160km/h 很高
}

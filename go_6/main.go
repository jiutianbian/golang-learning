package main

import (
	"fmt"
)

//课程：面向对象

//前面两章我们介绍了函数和struct，那你是否想过函数当作struct的字段一样来处理呢？
//今天我们就讲解一下函数的另一种形态，带有接收者的函数，我们称为method

//--------------------------示例一：method简单示例------------------------------//
//首先我们用struct来定义一个实体car，他有3个属性brand，price，speed
type car struct {
	brand string
	price string
	speed string
}

//其次如果我们定义一个show()方法，并在show方法前加上(c car)这个申明，那么这个show函数我们就称为是car这个实体的method
//这样面向对象所需要的基本属性和方法定义这样就都有了，这个就跟java里面的类，定义属性和方法很相似了
func (c car) show() {
	fmt.Println("我是:", c.brand, c.price, c.speed)
}

//--------------------------示例二：method 继承------------------------------//
//前面我们学习了字段的继承，那么你也会发现Go的一个神奇之处，method也是可以继承的。
//如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。让我们来看下面这个例子
type suvcar struct {
	car  // struct作为匿名字段
	size string
}

//--------------------------示例三：method 重写------------------------------//
//上面的例子中，如果suvcar想要实现自己的show,怎么办？简单，和匿名字段冲突一样的道理，
//我们可以在suvcar上面定义一个method，重写了匿名字段的方法。请看下面的例子
//func (c suvcar) show() {
//	fmt.Println("我是:", c.brand, c.price, c.speed, c.size)
//}

//指针作为接受者,比起直接传对象，可以用来保证设置属性的操作等真正作用于对象本身，而不是对象的copy
//func (c *car) setBrand(brand string) {
//	c.brand = brand
//}

//func (c *car) setPrice(price string) {
//	c.price = price
//}

//func (c car) setSpeed(speed string) {
//	c.speed = speed
//}

func main() {
	var baoma car
	baoma.brand = "宝马"
	baoma.price = "40万"
	baoma.speed = "120km/h"
	baoma.show()

	var suvbaoma suvcar
	suvbaoma = suvcar{car{"宝马suv", "80万", "160km/h"}, "很高"}

	// 实现了匿名继承的struct，也能继承他的method
	suvbaoma.show()
	// 如果没有重写show方法，输出  我是: 宝马suv 80万 160km/h
	// 如果重写了show方法，输出  我是: 宝马suv 80万 160km/h 很高

	// 指针作为接受者,比起直接传对象，可以用来保证设置属性的操作等真正作用于对象本身，而不是对象的copy
	//	benchi := new(car)
	//	benchi.setBrand("奔驰")
	//	benchi.setPrice("60万")
	//	benchi.setSpeed("140km/h")
	//	benchi.show()
	//输出 “我是: 奔驰 60万” 没有输出速度，因为setSpeed不是作用在car本生，而是作用在car的copy上
}

package main

import (
	"fmt"
)

//课程：结构体

//1.结构体（struct）是用户自定义的类型，它代表若干字段的集合。
//有些时候将多个数据看做一个整体要比单独使用这些数据更有意义，这种情况下就适合使用结构体。
//例如我们用type关键字来声明新的类型car，car是个结构体，里面拥有属性brand,price,speed 3个属性，
//这样的类型我们称之struct，代码如下：
type car struct {
	brand string
	price string
	speed string
}

//--------匿名字段--------
//2. struct中的属性是可以是struct类型的，且struct属性可以省略字段名作为匿名字段
//当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
//我们看下面的这个例子
type suvcar struct {
	car
	size string
}

func main() {
	//初始化
	//申明自定义类型car的一个变量baoma
	var baoma car

	baoma.brand = "宝马"
	baoma.price = "40万"
	baoma.speed = "120km/h"

	fmt.Println("我是：", baoma.brand, baoma.price, baoma.speed)

	//--------匿名字段代码演示--------//
	//有匿名字段的suvcar初始化,可以看到suvcar的属性，直接继承了car中的属性
	var suvbaoma suvcar
	suvbaoma.brand = "宝马"
	suvbaoma.price = "40万"
	suvbaoma.speed = "120km/h"
	suvbaoma.size = "我是suv，我很高"

	//可以直接访问匿名字段中的属性
	fmt.Println("我是：", suvbaoma.brand, suvbaoma.price, suvbaoma.speed, suvbaoma.size)

	//	//按照属性顺序顺序初始化
	//	var benchi car
	//	benchi = car{"奔驰", "60万", "140km/h"}
	//	fmt.Println("我是：", benchi.brand, benchi.price, benchi.speed)

	//	//键值对初始化
	//	var aodi car
	//	aodi = car{brand: "奥迪", price: "80万", speed: "160km/h"}
	//	fmt.Println("我是：", aodi.brand, aodi.price, aodi.speed)

	//	var saloonbaoma salooncar
	//	saloonbaoma = salooncar{car: car{"奔驰", "60万", "140km/h"}, color: "黑色", string: "很矮"}

	//	fmt.Println("我是：", saloonbaoma.brand, saloonbaoma.price, saloonbaoma.speed, saloonbaoma.color, saloonbaoma.string)
}

////自定义一个类型color
//type color string

////匿名字段除了可以是struct，也可以是自定义类型，内置类型
//type salooncar struct {
//	car    // struct作为匿名字段
//	color  // 自定义类型作为匿名字段
//	string // 内置类型作为匿名字段
//}

package main

import (
	"fmt"
)

//用type关键字来声明新的类型car，car是个结构体，里面拥有属性brand,price,speed
type car struct {
	brand string
	price string
	speed string
}

//struct中的属性是可以是struct类型的，且struct属性可以省略字段名作为匿名字段
type suvcar struct {
	car
	size string
}

//自定义一个类型color
type color string

//匿名字段除了可以是struct，也可以是自定义类型，内置类型
type salooncar struct {
	car    // struct作为匿名字段
	color  // 自定义类型作为匿名字段
	string // 内置类型作为匿名字段
}

func main() {
	//申明自定义类型car的一个变量baoma
	var baoma car
	//初始化
	baoma.brand = "宝马"
	baoma.price = "40万"
	baoma.speed = "120km/h"

	fmt.Println("我是：", baoma.brand, baoma.price, baoma.speed)

	//按照属性顺序顺序初始化
	var benchi car
	benchi = car{"奔驰", "60万", "140km/h"}
	fmt.Println("我是：", benchi.brand, benchi.price, benchi.speed)

	//键值对初始化
	var aodi car
	aodi = car{brand: "奥迪", price: "80万", speed: "160km/h"}
	fmt.Println("我是：", aodi.brand, aodi.price, aodi.speed)

	//指针方式初始化
	var kaidilake *car
	kaidilake = new(car)
	kaidilake.brand = "凯迪拉克"
	kaidilake.price = "100万"
	kaidilake.speed = "180km/h"
	fmt.Println("我是：", kaidilake.brand, kaidilake.price, kaidilake.speed)

	//有匿名字段的struct初始化,可以看到匿名字段中的属性，
	var suvbaoma suvcar
	suvbaoma.car = baoma
	suvbaoma.size = "很高"

	//可以直接访问匿名字段中的属性
	fmt.Println("我是：", suvbaoma.brand, suvbaoma.price, suvbaoma.speed, suvbaoma.size)

	var saloonbaoma salooncar
	saloonbaoma = salooncar{car: car{"奔驰", "60万", "140km/h"}, color: "黑色", string: "很矮"}

	fmt.Println("我是：", saloonbaoma.brand, saloonbaoma.price, saloonbaoma.speed, saloonbaoma.color, saloonbaoma.string)
}

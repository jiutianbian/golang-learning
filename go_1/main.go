package main

import (
	"errors"
	"fmt"
)

func main() {
	//变量申明与赋值分离
	var a string         //定义变量名为a的一个string字符串
	a = "nihao,wo shi a" //给变量a赋值

	//变量申明与赋值一起
	var b string = "nihao,wo shi b" //定义变量名为b的一个string字符串，直接赋值

	//多个变量申明与赋值
	var c, d string = "nihao,wo shi c", "nihao,wo shi d" //一起定义c和d两个字符串，并赋值

	//省略变量类型
	var e, f = "nihao,wo shi e", "nihao,wo shi f" //省略了变量类型，编译器会用相应值的类型来帮你初始化它们

	//省略变量类型和 var关键字
	g, h := "nihao,wo shi g", "nihao,wo shi h" //省略了var关键字和类型，通过:=来赋值，这边不能用＝号，编译器会根据初始化的值自动推导出相应的类型

	//定义常量
	const HELLO = "NIHAO"

	//定义Boolean型
	var isWoking bool = true

	//定义数值类型
	var i int = 1 //整数

	var j float32 = 6.3 //浮点数
	var k float64 = 5.4

	//定义复数类型
	var l complex64 = 5 + 5i

	//定义错误类型
	err := errors.New("chu cuo le,da xiong di")

	//定义 array变量
	var arr_1 [3]int = [3]int{1, 2, 3}
	arr_2 := [...]int{4, 5, 6}

	//定义 slice变量，与arry相比少了固定长度的申明
	var slice_1 []int = []int{1, 2, 3} //直接初始化
	slice_2 := arr_1[0:2]              //从数组中直接截取
	slice_3 := slice_2[0:1]            //从slice中直接截取

	// 声明一个key是字符串，值为字符串的字典,这种方式的声明需要在使用之前使用make初始化
	var map_1 map[string]string
	// 另一种map的声明方式
	map_1 = make(map[string]string)
	map_1["one"] = "one"
	map_1["two"] = "two"
	map_1["three"] = "three"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(HELLO)

	fmt.Println(isWoking)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Println(err)

	//_ 表示缺省赋值
	for _, i := range arr_1 {
		fmt.Println(i)
	}

	for _, i := range arr_2 {
		fmt.Println(i)
	}

	for _, i := range slice_1 {
		fmt.Println(i)
	}

	for _, i := range slice_2 {
		fmt.Println(i)
	}

	for _, i := range slice_3 {
		fmt.Println(i)
	}

	fmt.Println(map_1["one"])
	fmt.Println(map_1["two"])
	fmt.Println(map_1["three"])
}

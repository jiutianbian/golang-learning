package main

import (
	"errors"
	"fmt"
)

//课程：变量与常量的定义
func main() {
	//--------------------------示例一： 变量定义示例  ------------------------------//
	//go语言中使用关键字var来定义变量,可以有以下5种方式
	//其中1，2两种方式跟java定义类似，但是golang中：变量类型是申明在变量后面的，而Java是申明在变量前面的

	//其中3，4，5的申明方式与java相比，是go语言特有
	//其中第3种方式，go语言可以同时初始化多个变量
	//其中第4种方式，go语言省略了变量类型，编译器会用相应值的类型来帮你初始化它们
	//其中第5种方式：go语言省略了var关键字和类型，但是通过:=来赋值，这边不能用＝号，编译器会根据初始化的值自动推导出相应的类型
	//第5种方式是最简洁的申明方式，也是go语言中用的最多申明的方式，我们也可以看一bilibi后端源码

	//1.变量申明与赋值分离
	var a string         //定义变量名为a的一个string字符串，类型是写在变量后面的,这点跟java很不一样
	a = "nihao,wo shi a" //给变量a赋值

	//2.变量申明与赋值一起
	var b string = "nihao,wo shi b" //定义变量名为b的一个string字符串，直接赋值，java中的写法：String a="nihao";

	//3.多个变量申明与赋值
	var c, d string = "nihao,wo shi c", "nihao,wo shi d" //一起定义c和d两个字符串，并赋值  Java写法: String a="nihao",b="nibuhao";

	//4.省略变量类型
	var e, f = "nihao,wo shi e", "nihao,wo shi f" //省略了变量类型，编译器会用相应值的类型来帮你初始化它们，与java相比，这是go语言特有

	//5.省略变量类型和 var关键字，这个是实际开发中用的最多的变量申明方式
	g, h := "nihao,wo shi g", "nihao,wo shi h" //省略了var关键字和类型，通过:=来赋值，这边不能用＝号，编译器会根据初始化的值自动推导出相应的类型，与java相比，这是go语言特有

	//--------------------------示例二： 其他变量定义简介  ------------------------------//
	//基本类型
	//定义Boolean型
	var isWoking bool = true

	//定义数值类型
	var i int = 1 //整数

	var j float32 = 6.3 //浮点数，跟Java里float
	var k float64 = 5.4 //跟java比没有double

	//定义复数类型
	var l complex64 = 5 + 5i

	//定义常量，常量类型通过const关键字来申明
	const HELLO = "NIHAO"

	//定义错误类型，这个也是go语言中独有的，功能类似java里面的expection
	err := errors.New("chu cuo le,da xiong di")

	//arrry,slice,map
	//定义 array变量
	var arr_1 [3]int = [3]int{1, 2, 3}
	arr_2 := [...]int{4, 5, 6}

	//定义 slice变量，一个动态的数组，这个跟java里面的List差不多
	var slice_1 []int = []int{1, 2, 3} //直接初始化
	slice_2 := arr_1[0:2]              //从数组中直接截取
	slice_3 := slice_2[0:1]            //从slice中直接截取

	// map类型，这个跟java里面的MAP差不多
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
	for _, i := range arr_1 { //大括号不能换行放置
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

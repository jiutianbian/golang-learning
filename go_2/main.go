package main

import (
	"fmt"
)

func main() {
	/* defer语句 */
	//defer语句 用来预定对一个函数的调用。它只能出现在一个函数中（假设是A函数），且只能调用另一个函数（假设是B函数），意味着在A函数结束返回时，延迟调用B函数，一般用于打开文件时的资源清理等工作。如果一个函数内部调用多个 defer 语句，则遵循后进先出的原则。
	defer test_1()
	defer test_2()
	//输出：在main函数的最后执行，先输出nihao，再输出hahaha，后进先出。自我感觉，defer的功能有点java中finally的意思

	/* if语句 */
	x := 1
	if x < 3 {
		fmt.Println("x 小于 3")
	} else if x == 3 {
		fmt.Println("x 等于 3")
	} else {
		fmt.Println("x 大于 3")
	}

	//在if的条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方不起作用了
	if y := 3; y < 3 {
		fmt.Println("y 小于 3")
	} else if y == 3 {
		fmt.Println("y 等于 3")
	} else {
		fmt.Println("y 大于 3")
	}

	//	fmt.Println(y)

	/* for语句 */
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	//省略表达式，类似其他语言的while功能
	a := 0
	for a < 5 {
		fmt.Println("a:", a)
		a++
	}

	//运用break跳出循环
	for b := 0; b < 5; b++ {
		if b == 1 {
			break
		}

		fmt.Println("b:", b)
	}
	// break打印出来0

	//运用continue跳出本次循环
	for c := 0; c < 5; c++ {
		if c == 1 {
			continue
		}

		fmt.Println("c:", c)
	}
	// continue打印出来0,2,3,4

	//用for语句，来迭代string\切片\字典等类型
	m := map[string]int{"E": 1, "F": 2}
	for k, v := range m {
		fmt.Println("k:v=", k, v)
	}

	/* switch语句 */
	i := 1
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("你好，大兄弟")
	}

	//switch中，每个case是默认是自带break的，这与C语言不同，如果需要强制执行后面的语句,需要添加fallthrough
	switch i {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	default:
		fmt.Println("你好，大兄弟")
	}

	/* goto语句 */
	//用goto跳转到必须在当前函数内定义的标签，不建议实用goto，逻辑不太清晰，容易死循环
	d := 0
Here: //这行的第一个词，以冒号结束作为标签
	fmt.Println(d)
	d++

	if d < 5 {
		goto Here
	}
	//输出 0，1，2，3，4

}

func test_1() {
	fmt.Println("hahaha")
}
func test_2() {
	fmt.Println("nihao")
}

package main

import (
	"fmt"
)

//课程：流程控制
func main() {
	//--------------------------示例四： defer示例  ------------------------------//
	//defer是golang的一个特色功能，被称为“延迟调用函数”。
	//当外部函数返回后执行defer。
	//类似于java语言的 try… catch … finally… 中的finally，当然差别还是明显的。
	defer test_1()

	//--------------------------示例一： if示例  ------------------------------//
	//Go里面if条件判断语句中不需要括号，如下代码所示
	x := 1
	if x < 3 {
		fmt.Println("x 小于 3")
	} else if x == 3 {
		fmt.Println("x 等于 3")
	} else {
		fmt.Println("x 大于 3")
	}

	// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，如下所示
	// 用分号来隔离，不是逗号
	if y := 3; y < 3 {
		fmt.Println("y 小于 3")
	} else if y == 3 {
		fmt.Println("y 等于 3")
	} else {
		fmt.Println("y 大于 3")
	}

	//--------------------------示例二： for示例  ------------------------------//
	/* Go里面最强大的一个控制逻辑就是for，它既可以用来循环读取数据，又可以当作while来控制逻辑。它的语法如下： */
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	//for的省略表达式，go中没有while关键字的，省略表达式，就是类似java语言的while功能
	a := 0
	for a < 5 {
		fmt.Println("a:", a)
		a++
	}

	//--------------------------示例三： switch示例  ------------------------------//
	//switch中，每个case是默认是自带break的，这与java不同，如果需要强制执行后面的语句,需要添加fallthrough
	i := 1
	switch i {
	case 1:
		fmt.Println("没有fallthrough:1")
	case 2:
		fmt.Println("没有fallthrough:2")
	case 3:
		fmt.Println("没有fallthrough:3")
	default:
		fmt.Println("你好，大兄弟")
	}

	switch i {
	case 1:
		fmt.Println("有fallthrough：1")
		fallthrough
	case 2:
		fmt.Println("有fallthrough：2")
		fallthrough
	case 3:
		fmt.Println("有fallthrough：3")
		fallthrough
	default:
		fmt.Println("你好，大兄弟")
	}

	//用for语句，来迭代切片\字典等类型，
	//	m := map[string]int{"E": 1, "F": 2}
	//	for k, v := range m {
	//		fmt.Println("k:v=", k, v)
	//	}

	//break关键字
	//	for b := 0; b < 5; b++ {
	//		if b == 1 {
	//			break
	//		}

	//		fmt.Println("b:", b)
	//	}
	// break打印出来0

	//continue关键字
	//	for c := 0; c < 5; c++ {
	//		if c == 1 {
	//			continue
	//		}

	//		fmt.Println("c:", c)
	//	}
	// continue打印出来0,2,3,4

	/* goto关键字 */
	//用goto跳转到必须在当前函数内定义的标签，不建议使用goto，逻辑不太清晰，容易死循环
	//	d := 0
	//Here: //这行的第一个词，以冒号结束作为标签
	//	fmt.Println(d)
	//	d++

	//	if d < 5 {
	//		goto Here
	//	}
}

func test_1() {
	fmt.Println("hahaha")
}
func test_2() {
	fmt.Println("nihao")
}

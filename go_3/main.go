package main

import (
	"fmt"
)

//课程：函数定义

//--------------------------示例一： 函数简单定义  ------------------------------//
//1. golang中函数用 func 关键字来声明，sayHello为方法名
// (name string, age int)这串是入参，(sayName string, sayAge string, askYou string)是出参
//2. 函数入参，可以没有，或者单个，或者动态多个。
//3. 函数返回值，可以没有，单个或者多个返回值，这个要着重说一下，
//java和C最多只能有单个返回值，不存在多个返回值的情况，此属于golang函数的特性
//4. golang函数名，区分大小写，大写开头为公有方法，小写开头为私有，这个要着重说一下，
func sayHello(name string, age int) (sayName string, sayAge string, askYou string) {
	sayName = "我叫" + name                        //由于返回值中已经定义了sayName的变量，所以不需要再定义，可以直接赋值
	sayAge = "今年" + fmt.Sprintf("%d", age) + "岁" //将int类型通过Sprintf方法转为字符串
	askYou = "你呢？"

	return sayName, sayAge, askYou
}

//--------------------------示例二： 变餐  ------------------------------//
//golang中函数支持变参。接受变参的函数是有着不定数量的参数的。变参是通过 ... 加上变量类型，组合起来申明的
func saySomeThing(something ...string) {
	for index, value := range something {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}
}

//--------------------------示例三： 指针做为参数  ------------------------------//
//golang中还可以传入指针做为参数，什么是指针，我们下个知识点再带大家了解，这个指针做为参数的，大家就先了解一下就好
func sayHelloByNameZhiZhen(name *string) (sayHelloByName string) {
	*name = *name + "汉克斯"
	return "你好" + *name
}

//--------------------------示例四： 函数做为参数  ------------------------------//
//在golang中函数也是一种变量，我们可以通过type来定义它
//例如 我们定义了一个函数类型为sayLoveOrHate，它用来表示：一个输入string参数和一个输出string参数的所有函数
type sayLoveOrHate func(string) string

//例如我们定义的sayLove函数和sayHate函数，它们就是sayLoveOrHate类型
func sayLove(name string) (sayLove string) {
	return "我喜欢" + name
}

func sayHate(name string) (sayLove string) {
	return "我恨" + name
}

//然后将sayLoveOrHate做为参数传入到函数中
//例如声明了一个函数getMyView，它传入了sayLoveOrHate的函数类型作为输入参数
func getMyView(name string, say sayLoveOrHate) (view string) {
	return say(name)
}

func main() {

	sayName, sayAge, askYou := sayHello("小明", 12)
	fmt.Println(sayName)
	fmt.Println(sayAge)
	fmt.Println(askYou)

	saySomeThing("你好", "能给下电话号码么", "看在天气这么好的份儿上")

	name := "汤姆"

	//如果我们只是传入值，不传指针，那么只相当于对name的copy做操作，实际的name值不会变更，还是会输出汤姆
	sayHelloByName := sayHelloByName(name)
	fmt.Println("name:" + name) //输出 name:汤姆

	//如果我们传指针，那么只相当于对name本身做操作，实际的name值会变更，会输出为汤姆汉克斯
	sayHelloByNameZhiZhen := sayHelloByNameZhiZhen(&name)
	fmt.Println("namebyzhizhen:" + name) //输出 namebyzhizhen:汤姆汉克斯

	fmt.Println("sayHelloByName:" + sayHelloByName)               //输出 sayHelloByName:你好汤姆汉克斯
	fmt.Println("sayHelloByNameZhiZhen:" + sayHelloByNameZhiZhen) //输出 sayHelloByNameZhiZhen:你好汤姆汉克斯

	//通过传入同种输入输出格式的函数，但不同的函数，来获得不同的函数输出，这样子使我们的程序变得非常的灵活
	fmt.Println(getMyView("小明", sayHate))    //输出 我恨小明
	fmt.Println(getMyView("汤姆汉克斯", sayLove)) //输出 我喜欢汤姆汉克斯

	//	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
	//		if err := recover(); err != nil {
	//			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	//		}
	//	}()

	//	panic("我死掉了") //主动抛出 panic
}

func sayHelloByName(name string) (sayHelloByName string) {
	name = name + "汉克斯"
	return "你好" + name
}

package main

import (
	"fmt"
)

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

	//通过传入同种输入输出格式的函数，但不同的函数，来获得不同的函数输出
	fmt.Println(getMyView("小明", sayHate))    //输出 我恨小明
	fmt.Println(getMyView("汤姆汉克斯", sayLove)) //输出 我喜欢汤姆汉克斯

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()

	panic("我死掉了") //主动抛出 panic
}

func sayHello(name string, age int) (sayName string, sayAge string, askYou string) {
	sayName = "我叫" + name                        //由于返回值中已经定义了sayName的变量，所以不需要再定义，可以直接赋值
	sayAge = "今年" + fmt.Sprintf("%d", age) + "岁" //将int类型通过Sprintf方法转为字符串
	askYou = "你呢？"

	return sayName, sayAge, askYou
}

func saySomeThing(something ...string) { //变参是指不定数量的参数，用...和具体的参数类型表示
	for index, value := range something {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}
}

func sayHelloByName(name string) (sayHelloByName string) {
	name = name + "汉克斯"
	return "你好" + name
}

func sayHelloByNameZhiZhen(name *string) (sayHelloByName string) {
	*name = *name + "汉克斯"
	return "你好" + *name
}

type sayLoveOrHate func(string) string // 声明了一个函数类型为sayLoveOrHate，用来表示一个输入参数和一个输出参数的函数

func getMyView(name string, say sayLoveOrHate) (view string) { //声明了一个函数getMyview，传入了sayLoveOrHate的函数类型作为输入参数
	return say(name)
}

func sayLove(name string) (sayLove string) {
	return "我喜欢" + name
}

func sayHate(name string) (sayLove string) {
	return "我恨" + name
}

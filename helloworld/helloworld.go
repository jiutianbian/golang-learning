package main //go语言中包（package）与java中的包（package）非常类似，都是组织代码的方式
//go语言中，包名一般为go代码所在的目录名，但是不绝对；与java不同的是，go语言中包名只有一级，而在java中包名是以点分割的多级目录组合的。

import "fmt" //go语言中，通过import关键字不是include关键字来导入引用包, import "fmt"：这边表示引入了系统打印功能的fmt包
//这个也与java中的improt类似，但go中不能引入无用的包

func main() { //go语言中，通过func关键字来申明函数方法，此为main函数方法，系统默认的生命周期方法，程序启动会自动被执行,是程序启动入口，跟java里面的main方法类似
	fmt.Println("hello呀～～～～～～～～～～～～～～～～～～～～") //打印出消息
}

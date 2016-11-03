package main

import (
	"fmt"
)

func main() {
	var name string

	//定义了一个指针p,用*加具体类型表示，用来表示变量的地址
	var p *string

	name = "tom"

	//将name的地址赋值给指针p,用&取地址
	p = &name

	fmt.Println(&name)
	fmt.Println(p)
	fmt.Println(*p) //用*表示取指针指向地址的值

	//打印
	//0xc82000a290
	//0xc82000a290
	//tom

	//go中 还可以通过new命令分配内存地址，返回指针类型，如下
	var p1 *int

	p1 = new(int)

	fmt.Println(p1)
	fmt.Println(*p1)
	//打印
	//0xc82000a2c0
	//0
}

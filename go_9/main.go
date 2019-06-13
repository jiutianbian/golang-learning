package main

import (
	"fmt"
	"net/http"
)

//课程：简单web服务器
//Go语言里面提供了一个完善的net/http包，
//通过http包可以很方便的搭建起来一个可以运行的Web服务。
//同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。

func goodgoodstudy(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)        //通过 request，执行http请求的相关操作
	response.Write([]byte("day day up")) //通过 response，执行http的响应相关操作
}

func nihao(response http.ResponseWriter, request *http.Request) {
	fmt.Println("nihao~~~")
	response.Write([]byte("ni hao"))
}

func main() {
	http.HandleFunc("/", goodgoodstudy) //设置访问的监听路径，以及处理方法
	http.ListenAndServe(":9000", nil)   //设置监听的端口

	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/", goodgoodstudy) //设置访问的监听路径，以及处理方法
	//	mux.HandleFunc("/nihao", nihao)

	//	http.ListenAndServe(":9000", mux) //设置监听的端口
}

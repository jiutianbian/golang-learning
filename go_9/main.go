package main

import (
	"fmt"
	"net/http"
)

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

	mux := http.NewServeMux()
	mux.HandleFunc("/", goodgoodstudy) //设置访问的监听路径，以及处理方法x
	mux.HandleFunc("/nihao", nihao)

	http.ListenAndServe(":9000", mux) //设置监听的端口
}

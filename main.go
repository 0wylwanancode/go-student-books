package main

import (
	"net/http"

	"stumanage/pkg/routers"
)

func main() {
	mx := http.NewServeMux() //返回*http.ServeMux
	//
	//

	//调用路由
	//启动路由
	routers.RegisterStudentRouter(mx)

	http.ListenAndServe(":8990", mx)
	/*
		func ListenAndServe(addr string, handler Handler)
		Handler类型是一个实现了 ServeHTTP(ResponseWriter, *Request)方法的接口
		而*http.ServeMux正好实现了ServeHTTP(ResponseWriter, *Request)	ServeMux) ServeHTTP
		所以能当成参数传递过去,
		由于mx是新new()出来的结构体内未有值所以=nil

	*/
}

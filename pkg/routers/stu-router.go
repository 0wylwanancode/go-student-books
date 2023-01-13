package routers

import (
	"net/http"
	"stumanage/pkg/contor"
)

//增删改查的路由函数
var RegisterStudentRouter = func(mx *http.ServeMux) {
	/* 	http.HandleFunc() == mx.HandleFunc()
		原因
		http.HandleFunc()是去调用	DefaultServeMux.HandleFunc(pattern, handler)\
		var DefaultServeMuxs = &ServeMux{}
		说明DefaultServeMux是*http.ServeMux类型
		说明
		http.HandleFunc()=*http.ServeMux.HandleFunc() 而mx=*http.ServeMux
		所以 http.HandleFunc() == mx.HandleFunc()
	}
	*/
	mx.HandleFunc("/create", contor.Create) //创建
	mx.HandleFunc("/delete/", contor.Delete)
	mx.HandleFunc("/update/", contor.Update)
	//路径索引的话/fetch/xx/xx都能触发这个路由的视图函数
	mx.HandleFunc("/fetch/", contor.Fetch)
	mx.HandleFunc("/retrieve", contor.Retrieve)

}

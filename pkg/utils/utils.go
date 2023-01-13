package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//查询用户信息 获得用户i索引
func ParesParams(r *http.Request) (int, error) {
	path := r.URL.Path
	//
	pathSlice := strings.Split(path, "/")
	// fmt.Println(pathSlice, len(pathSlice)) //实际上我们要获切片最后一个内容

	id := pathSlice[len(pathSlice)-1] //获得切片最后一个内容
	//id为字符串为了匹配结构体的Int
	//  t=>i  字符串转化成int
	return strconv.Atoi(id)
}

//获取前端 body中json格式传来的数据

func ParesBody(r *http.Request, v interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		return json.Unmarshal(body, v)
	} else {
		return err
	}

}

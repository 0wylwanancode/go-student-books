package contor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"stumanage/pkg/models"
	"stumanage/pkg/utils"
)

var students []models.Student

//main前执行
func init() {
	//先创建一些数据
	for i := 0; i < 5; i++ {
		student := models.Student{
			Id:      i + 1,
			Name:    fmt.Sprintf("wang%d", i),
			Class:   fmt.Sprintf("三年级%d班", i),
			Chinese: 88,
			Math:    90,
			English: 69,
		}
		//将student数据循环填充到students
		students = append(students, student)
	}

}

//增加
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//POST body中json中传过来的值是个创建结构体中的Json定义的名称匹配
	//如果前端传递过来json,在结构体定义的json 对应key值, 则结构体的key为默认值

	Userdata := make([]byte, r.ContentLength) //创建一个长度为r.ContentLength
	user := models.Student{}
	r.Body.Read(Userdata)           //将r.body的值写入Userdata
	json.Unmarshal(Userdata, &user) //再将切片转化成,到第二个参数中
	fmt.Println(user)
	//新的学生信息id永远比最后一个学生信息id+1
	user.Id = students[len(students)-1].Id + 1
	students = append(students, user)

	//读取前端POST请求中body中json的数据传递
	/* strings := models.Student{}
	body, err := ioutil.ReadAll(r.Body)//将r.body中的值读取到body中去
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(body, &strings)
	fmt.Println(strings) */
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user)
	w.Write(res)
}

//删除
func Delete(w http.ResponseWriter, r *http.Request) {
	//

	if strings.ToUpper(r.Method) != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//获得前端传递的索引
	id, _ := utils.ParesParams(r)
	index := -1
	for i := 0; i < len(students); i++ {
		if students[i].Id == id {
			fmt.Fprint(w, "成功删除：")
			s := students[i]
			//将s ,json byte化
			DeleteUser, _ := json.Marshal(s)
			w.Header().Set("Context-Type", "application/json")
			w.Write(DeleteUser)
			fmt.Fprint(w, "学生信息")
			//         0 1 2 3 4
			//我们要删除的是索引为i的切片内容 students[:i]获得切片 是不包含i下标这个值
			students = append(students[:i], students[i+1:]...)
			//赋值删除的学生id
			index = id
			w.WriteHeader(http.StatusOK)
			break //这个break可以跳出 if之外的for循环
		}

	}
	//说明没有找到删除的学生信息
	if index == -1 {
		fmt.Fprintf(w, "不存在%d学号学生,无法进行删除操作", id)
		w.WriteHeader(http.StatusOK)
		return
	}

}

//改
func Update(w http.ResponseWriter, r *http.Request) {
	//判断请求方式是否是PUT
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id, _ := utils.ParesParams(r)

	index := -1
	//找到对应学号学生
	for i, _ := range students {
		//当有当有学生学号匹配到穿过的值时,将学号保存
		if students[i].Id == id {
			//记录学生学号
			index = i
		}

	}
	if index == -1 {
		fmt.Fprintf(w, "未找到：%d学号学生", id)
		return
	}
	//获取前端PUT请求
	//获得用户信息
	stu := models.Student{}
	//调用函数传过去的一定要是内存,不然无法修改stu的值
	err := utils.ParesBody(r, &stu)
	if err != nil {
		panic(err)
	}

	//当获取的前端传递的Chinese不为0
	if stu.Chinese != 0 {
		students[index].Chinese = stu.Chinese
	}
	//当获取的前端传递的Math不为0
	if stu.Math != 0 {
		students[index].Math = stu.Math
	}
	//当获取的前端传递的English不为0
	if stu.English != 0 {
		students[index].English = stu.English
	}
	fmt.Println("修改后的学生信息", students[index])

	//修改返回的数据类型
	w.Header().Set("Context-Type", "application/json")

	//再将修改后的数据返回给前端
	res, err := json.Marshal(students[index])
	if err != nil {
		fmt.Println(err)
		return
	}

	//设置返回响应的值
	w.WriteHeader(http.StatusOK)
	//将序列化的信息返回给前端
	w.Write(res)

}

//获取某个学生信息 只接收Get请求
func Fetch(w http.ResponseWriter, r *http.Request) {
	//如果不是GET请求直接返回405
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	///fetch/1获得查询的用户id
	idint, _ := utils.ParesParams(r)
	index := -1
	for i := 0; i < len(students); i++ {
		//当查询的用户id==某个用户的id时,返回这个用户的信息
		if students[i].Id == idint {
			//结构体格式化成[]byte byte 为int8为 ASII
			b, _ := json.Marshal(students[i])
			w.Header().Set("Context-Type", "application/json")
			w.Write(b)
			//返回200
			//配合到了赋值给Index
			index = i
			return
		}
	}
	if index == -1 {
		fmt.Fprintf(w, "没有学号为：'%d'号这名学生", idint)
		w.WriteHeader(http.StatusOK)
	}

}

//查询全部学生信息,只接收Get请求
func Retrieve(w http.ResponseWriter, r *http.Request) {
	//设置访问类型
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Context-Type", "application/json")
	// fmt.Fprintf(w, "%v", students)
	b, _ := json.Marshal(students)
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

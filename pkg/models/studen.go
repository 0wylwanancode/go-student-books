package models

type Student struct {
	Id      int     `json:"id" form:"id"`           //学号
	Name    string  `json:"name" form:"name"`       //姓名
	Class   string  `json:"class" form:"class"`     //班级
	Chinese float32 `json:"chinese" form:"chinese"` //语文
	Math    float32 `json:"math" form:"math"`       //数学
	English float32 `json:"english" form:"english"` //英文
}

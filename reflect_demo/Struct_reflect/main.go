package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type student struct {
	Name  string `json:"name" xiaokeai:"linyingmei"`
	Score int    `json:"score" xiaokeai:"bijie"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 18,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	//	 通过for 循环遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("xiaokeai"))
	}
	//	通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Println(scoreField.Name)
	}
}

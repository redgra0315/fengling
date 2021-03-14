package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Int8:
		fmt.Printf("type is int,value is %d\n", int8(v.Int()))
	case reflect.String:
		fmt.Printf("type is string,value is %v\n", string(v.String()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %v\n", float32(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int8 {
		// 尝试修改副本的值，引发panic
		v.SetInt(200)
	}
}

// 通过反射设置变量的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		// 尝试修改副本的值，引发panic
		v.Elem().SetInt(200)
	}
}
func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	c := "hello world"
	reflectType(b)

	var C = Cat{}
	reflectType(C)

	reflectValue(a)
	reflectValue(b)
	reflectValue(c)

	// 设置值
	reflectSetValue2(&b)
	fmt.Print(b)
}

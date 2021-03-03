package main

import (
	"fmt"
	"reflect"
)

func demo() {
	var x1 float64 = 3.4
	fmt.Printf("type:%v value:%v\n", reflect.TypeOf(x1), reflect.ValueOf(x1))

	var x2 float64 = 3.4
	v1 := reflect.ValueOf(x2)
	fmt.Println("type:", v1.Type())
	fmt.Println("kind is float64:", v1.Kind() == reflect.Float64)
	fmt.Println("value:", v1.Float())

	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())

	fmt.Println("_________")
	type MyInt int
	var x3 MyInt = 7
	v3 := reflect.ValueOf(x3)
	fmt.Printf("type: %v, value: %v\n", reflect.TypeOf(v3), reflect.ValueOf(v3))
	fmt.Println(v3)

	y := v1.Interface().(float64)
	fmt.Println(y)
}
func demo1(v float64) {
	var z float64 = 3.4
	//xz := z.Elem()

	p := reflect.ValueOf(&z) // Note: take the address of x.
	//p.SetFloat(v)
	xz := p.Elem()

	xz.SetFloat(v)
	fmt.Println(xz.Interface())
	fmt.Println(xz)
	//fmt.Println(p)
	fmt.Println("type of p:", xz.Type())
	fmt.Println("settability of p:", xz.CanSet())
}

func demo2() {
	type T struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func main() {
	//demo()

	//demo(10.2)
	demo2()
}

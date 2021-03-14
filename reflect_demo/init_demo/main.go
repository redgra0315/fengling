package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 解析config.ini 配置文件

type Mysql struct {
	Address  string `ini:"address"`
	Dbname   string `ini:"dbname"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
}

type Redis struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
}
type Config struct {
	Mysql `ini:"mysql"`
	Redis `ini:"redis"`
}

func loadInt(fileName string, data interface{}) (err error) {
	// 0. 参数校验
	// 0.1.传进来的data参数必须是指针类型
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		// 格式化输出一个error类型
		err = fmt.Errorf("类型错误")
		return
	}

	// 0.2.传进来的data参数必须是结构体类型
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("类型错误  struct..")
		return
	}
	// 1.读取配置文件
	o, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
		return
	}
	//string(o) // 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(o), "\n")
	fmt.Println(lineSlice)
	// 2.一行一行读
	var structName string
	for index, line := range lineSlice {
		// 去掉空格
		line = strings.TrimSpace(line)
		// 去掉空行
		if len(line) == 0 {
			continue
		}
		// 2.1.如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头的就表示是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax err", index+1)
				return
			}
			// 把这一行的首尾的[]去掉，取道中间的内容把首尾的空格去掉拿取里面的内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax err", index+1)
				return
			}
			// 根据字符串sectionName 去data里面查找反射对应的结构体
			//value := reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 记录字段名
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)

				}
			}
		} else {
			//2.3 如果不是[开头就是=分割的键值对
			//2.4 以=分割符号为一行，等号左边是key,右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d stntax error", index+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//  2.5 根据structname 去data里面把对应的嵌套结构体取出
			v := reflect.ValueOf(data)
			// 拿到结构体的值信息
			sValue := v.Elem().FieldByName(structName)
			// 拿到嵌套结构体的类型
			sType := sValue.Type()
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			//	遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				// TAG 信息是存储在类型信息中的
				filed := sType.Field(i)
				fileType = filed
				if filed.Tag.Get("ini") == key {
					//	找到对应字段名
					fieldName = filed.Name
					break
				}
			}
			//	4.如果key=tag,给这个字段赋值
			//	4.1 根据filedName取出字段
			if len(fieldName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fileName)
			// 4.2 对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fmt.Println(value)
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}

		}
	}
	return nil
}

func main() {
	var mc Config
	//var x = new(int)
	err := loadInt("reflect_demo/init_demo/config.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}
	//fmt.Println(mc.Address, mc.Port)
}

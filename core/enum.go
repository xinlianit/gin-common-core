package core

import (
	"reflect"
)

type EnumType int

type Enum struct {
	Value int
	Name  string
}

func New(enum EnumType) Enum {
	name := ""
	// 获取结构体反射
	refEnum := reflect.ValueOf(enum)
	// 获取反射方法
	method := refEnum.MethodByName("Name")

	// 验证反射方法是否有效
	if valid := method.IsValid(); valid {
		// 调用反射方法
		values := method.Call(make([]reflect.Value, 0))
		// 获取方法返回值
		name = values[0].String()
	}

	return Enum{Value: int(enum), Name: name}
}

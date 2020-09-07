package core

// 枚举类型
type EnumType int

// 枚举名称
var EnumNames map[EnumType]string

// 包初始化
func init() {
	EnumNames = make(map[EnumType]string)
}

// 枚举值
func (e EnumType) Value() int {
	return int(e)
}

// 枚举名称
func (e EnumType) Name() string {
	return EnumNames[e]
}

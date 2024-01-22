package model

type User struct {
	Name      string
	Age       int
	Height    float32
	Education string
	Hobby     []string
	MoreInfo  map[string]interface{}
	// MoreInfo是一个字典 map.
	// 其中 key 是 string, 而 value是 interface{} 类型
	// interface{} 是 Go 中的一个特殊类型，表示空接口 empty interface, 能够存储具有字符串键的任意类型的值
}

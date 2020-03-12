package interview

import "fmt"


//代码输出什么
func InterfaceMain()  {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

//解析： 当且仅当接口的动态值和动态类型为nil时，接口类型值才为nil



//代码输出什么
func MakeMain()  {
	s := make(map[string]int)
	delete(s,"h")
	fmt.Println(s["h"])
}

//解析：删除map不存在的值不会报错，相当于没有任何操作。获取不存在的值时，返回类型对应的零值。
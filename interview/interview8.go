package interview

import "fmt"

func hello() []string  {
	return nil
}

//这个函数输出什么，原因
func helloMain()  {
	h := hello
	if h == nil {
		fmt.Println("nil")
	}else {
		fmt.Println("not nil")
	}
}

func GetValue() int  {
	return 1
}

//解析： not nil， 是将hello赋值给变量h，而不是函数的返回值，所以输出not nil



//以下函数能否编译通过，为什么
/*func GetValueMain()  {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}*/

//解析：编译失败，只有接口类型才可以使用类型选择。
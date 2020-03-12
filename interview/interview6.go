package interview

import "fmt"

type MyInt1 int
type MyInt2 = int

func IntTest()  {
	/*var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1,i2)*/
	fmt.Println("")
}

//解析：编译不通过，cannot use i（type int）as type MyInt1 in assignment
//MyInt1 int是基于int创建新类型MyInt
//MyInt2 = int 是定义了int的别名

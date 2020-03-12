package interview

import "fmt"

type person struct{
	name string
}

//这个函数输出什么：
func personMain()  {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

//解析：打印一个map中不存在的值时，返回元素类型的0值。




func hello1(num ...int)  {
	num[0] = 18
}

//这个函数输出什么：
func Hello1Main()  {
	i := []int{5,6,7}
	hello1(i...)
	fmt.Println(i[0])
}

//解析： 18
package interview

import "fmt"

//代码的运行结果
func struceTest()  {
	sn1 := struct {
		age int
		name string
	}{age:11,name:"qq"}
	sn2 := struct {
		age int
		name string
	}{age:11,name:"qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age:11,m: map[string]string{"a":"1"}}
	sm2 := struct {
		age int
		m map[string]string
	}{age:11,m: map[string]string{"a":"1"}}

	/*if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}*/
	fmt.Println(sm1,sm2)
}

//编译不通过 invalid operation：sm1 == sm2
//结构体只能比较是否相等，不能比较大小
//结构体类型相同才可以比较，结构体是否相同不但与属性有关，还与属性顺序有关
//map属于不可比较类型
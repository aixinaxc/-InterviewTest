package main

import "fmt"

func main()  {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)
	for key,val := range slice{
		m[key] = &val
	}

	for k,v := range m{
		fmt.Println(k,"->",*v)
	}
}


//解析: for range循环的时候会创建每个元素的副本，而不是元素引用，所以m[key]=&val取得都是变量val的地址，所以最后
//map中的所有元素的值变量val的地址，所以最后都是3.
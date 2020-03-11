package interview

import "fmt"

func AppendFunc() {
	s := make([]int,5)
	s = append(s,1,2,3)
	fmt.Println(s)
}

func AppendFunc1() {
	s := make([]int,0)
	s = append(s,1,2,3)
	fmt.Println(s)
}

//解析：make创建指定大小的切片后，会初始化值。同时append会像切片之后追加值。
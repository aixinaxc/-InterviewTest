package interview

import "fmt"


//这个函数输出什么：
func TypeAddMain()  {
	a := 5
	b := 8.1
	fmt.Println(a+b)
}

//解析：Invalid operation: a+b (mismatched types int and float64) ，类型不同的数值不能相加


//这个函数输出什么：
func IntMain()  {
	a := [5]int{1,2,3,4,5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

//解析：操作符 [i,j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，从索引 i，到索引 j 结束，
// 截取已有数组（切片）的任意部分，返回新的切片，新切片的值包含原数组（切片）的 i 索引的值，
// 但是不包含 j 索引的值。i、j 都是可选的，i 如果省略，默认是 0，j 如果省略，默认是原数组（切片）的长度。i、j 都不能超过这个长度值。
//假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：长度：j-i，容量：k-i。
//截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，
// 但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。
//所以例子中，切片 t 为 [4]，长度和容量都是 1。



func CompareMain()  {
	a := [2]int{5,6}
	b := [3]int{5,6}
	if a == b {
		fmt.Println("equal")
	}else {
		fmt.Println("not equal")
	}
}

//解析：Invalid operation: a == b (mismatched types [2]int and [3]int)
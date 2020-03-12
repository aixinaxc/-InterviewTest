package interview

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

//代码输出什么
func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

//解析： 0 1  。

package interview

import "fmt"

func mainA() {
	a := 1
	for i := 0; i < 5; i++ {
		a := a + 1
		a = a * 2
	}
	fmt.Print(a)
}

//解析：1  变量作用域。
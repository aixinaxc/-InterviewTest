package interview

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

	}()

	for {
		fmt.Print()
	}
}

//解析：runtime.GOMAXPROCS(1)指定只使用一个CPU，for循环会导致其他协程无法执行
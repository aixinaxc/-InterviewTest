package interview

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	// A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}


//解析：程序抛异常。先定义下，第一个协程为 A 协程，第二个协程为 B 协程；当 A 协程还没起时，
// 主协程已经将 channel 关闭了，当 A 协程往关闭的 channel 发送数据时会 panic，panic: send on closed channel。
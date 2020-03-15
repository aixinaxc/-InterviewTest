package interview

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var ints = make([]int, 0, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(len(ints))
}


//输出值不定。append()并不是并发安全的。
package interview

import "fmt"

func DeferFull()  {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()
	panic("触发异常")
}
//解析: defer的执行顺序是后进先出，当出现panic语句的时候，会先按照defer的后进先出的顺序执行，最后执行panic





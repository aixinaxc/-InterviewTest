package interview

import (
	"fmt"
	"io/ioutil"
	"os"
)

//下面代码存在什么问题
func main() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}

//解析：f可能为nil，如果defer放到判断前面，可能导致f.Close()报错

func fd() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func mainFD() {
	fd()
}

//recover:1。知识点：panic、recover()。当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容。

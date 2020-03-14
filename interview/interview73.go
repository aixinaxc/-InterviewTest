package interview

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

//解析：0xc000018058 2
//0xc000018058 2
//知识点：闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。（说明：i 的地址，输出可能与上面的不一样）。




var ff = func(i int) {
	print("x")
}

func mainF() {
	ff := func(i int) {
		print(i)
		if i > 0 {
			ff(i - 1)
		}
	}
	ff(10)
}


//解析：10x。这道题一眼看上去会输出 109876543210，其实这是错误的答案，
// 这里不是递归。假设 main() 函数里为 f2()，外面的为 f1()，当声明 f2() 时，调用的是已经完成声明的 f1()。
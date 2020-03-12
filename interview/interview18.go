package interview


//f1\f2\f3分别输出什么？
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//解析:f2调用时，先给局部变量t赋值5，然后将t赋值给返回值r，defer函数修改的是局部变量t，返回值r并没有被修改，故结果为5；
//f3调用时，先给返回值赋值为1，然后defer函数传进去的r是原来r的copy，不会改变要返回的那个r值，故结果是1
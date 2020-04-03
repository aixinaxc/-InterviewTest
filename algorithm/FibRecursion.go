package algorithm


//斐波纳契数（Fibonacci sequence），又称黄金分割数列，因数学家列昂纳多·
//斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，故又称为“兔子数列”，
//指的是这样一个数列：0、1、1、2、3、5、8、13、21、34、……
//用文字来说，就是斐波那契数列由0和1开始，之后的斐波那契数就是由之前的两数相加而得出。
//在数学上，斐波纳契数列以如下被以递归的方法定义：F(0)=0，F(1)=1, F(n)=F(n-1)+F(n-2)（n>=2，n∈N*）
func FibRecursion(n int) int {
	switch {
	case n < 2:
		return n
	default:
		return FibRecursion(n-1) + FibRecursion(n-2)
	}
}

func FibArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		switch {
		case i < 2:
			arr[i] = i
		default:
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}

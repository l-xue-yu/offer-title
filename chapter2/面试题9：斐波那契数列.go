package main

/*
知识点1：
什么是斐波那契数列？0、1、1、2、3、5、8、13、21、34……这个数列从第3项开始，每一项都等于前两项之和。
知识点2：
递归的原理，优势，劣势。
劣势：
递归由于是函数调用自身，而函数调用是有时间和空间的消耗的：每一次函数调用，都需要在内存栈中分配空间以保存参数、返回地址及临时变量，而且往栈里压入数据和弹出数据都需要时间。
优势：
代码要简洁很多，更加容易实现。
*/

func main() {
	println("斐波那契数列递归：", fibonacci(8))
	println("斐波那契数列循环：", fibonacciCycle(8))
}

/*
低效递归
重复计算的数据非常多，保存的栈，内存信息也多
*/
func fibonacci(n uint) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

/*
用循环实现，避免计算重复数据，和额外的内存信息
思路：从底到高，一级一级计算
*/
func fibonacciCycle(n int) int {

	result := []int{0, 1}
	if n < 2 {
		return result[n]
	}

	fibPre1 := result[0]
	fibPre2 := result[1]
	fibData := 0
	for i := 2; i <= n; i++ {
		fibData = fibPre1 + fibPre2
		fibPre1 = fibPre2
		fibPre2 = fibData
	}
	return fibData
}

package main

import "fmt"

/*
题目：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列1、2、3、4、5是某栈的压栈序列，序列4、5、3、2、1是该压栈序列对应的一个弹出序列，但4、3、5、1、2就不可能是该压栈序列的弹出序列。
*/

/*
思路：
感觉是不是我理解简单了，就是判断两个切片是否是相反就可以了
*/

/*
遍历入栈切片，如果不相等就退出
*/
//func validateStackSequences(pushed []int, popped []int) bool {
//
//	//异常数据返回
//	if len(pushed) != len(popped) || len(pushed) == 0 || len(popped) == 0 {
//		return false
//	}
//
//	pushLen := len(pushed)
//
//	for _, value := range pushed {
//		if value != popped[pushLen-1] {
//			return false
//		}
//		pushLen--
//	}
//	return true
//}

func main() {
	pushSlice := []int{1, 2, 3, 4, 5}
	popedSlice := []int{5, 4, 3, 2, 1}
	res := validateStackSequences(pushSlice, popedSlice)
	fmt.Println(res)
}

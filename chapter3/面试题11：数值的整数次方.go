package chapter3

import "fmt"

/*
题目：实现函数double Power（double base, int exponent），求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
*/

func main() {
	fmt.Printf("%f", powerWithUnsignedExponent(2, 6))
	//println(powerWithUnsignedExponent(2,6))
}

func Power(base float64, expoent int) float64 {
	if base == 0 { //0的情况
		return 0
	}
	if expoent == 0 {
		return 1
	}
	if expoent == 1 {
		return base
	}
	isLessZero := false
	if base < 0 { //负数的情况
		base = base * -1
		isLessZero = true
	}

	var result float64
	result = 1.0
	for i := 1; i < expoent; i++ {
		result *= base
	}
	if isLessZero {
		result = result * -1
	}
	return result
}

/*
使用右移代替%2
优化写法：使用公式
an={an/2*an/2 n为偶数
a(n-1)/2.a(n-1)/2.a n为奇数
递归，

*/

func powerWithUnsignedExponent(base float64, expoent int) float64 {
	if base == 0 { //0的情况
		return 0
	}
	if expoent == 0 {
		return 1
	}
	if expoent == 1 {
		return base
	}
	result := powerWithUnsignedExponent(base, expoent>>1)
	result *= result
	if expoent&1 == 1 {
		result *= base
	}
	return result
}

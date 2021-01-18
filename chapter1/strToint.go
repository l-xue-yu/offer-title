package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := strToInt("123456")
	if err != nil {
		println(err.Error())
		return
	}
	println(res)
}

/*
题目：将字符串中的int数据转换成int整数后返回
思路：遍历字符串，
每个字符就是一个uint8，而uint8又代表了ascii中的十进制，所以可以使用ascii的思路来直接替换每个字符，再根据每个字符的位置*10的n次方，
然后每个字符都累加起来。
入参：字符串
返回：int
知识点：
1.设计一个程序要考虑异常情况，比如这里的字符中可能有英文字母或其他字符，应该排除
2.golang 中byte uint8 ascii的互换关系
*/
func strToInt(strData string) (uint32, error) {
	//1.有非法字符，返回错误
	for i := len(strData) - 1; i >= 0; i-- {
		if strData[i] < 48 || strData[i] > 57 {
			return 0, errors.New("非法字符")
		}
	}

	var resInt uint32
	var i int
	var leng uint8
	leng = uint8(len(strData))
	ii := 0
	for i = int(leng - 1); i >= 0; i-- {
		resInt += factorial(i, strData[ii]-48)
		ii++
	}
	return resInt, nil
}

//10阶乘
func factorial(n int, val uint8) uint32 {
	var index int
	var facVal uint32
	facVal = 1
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for index = 1; index <= n; index++ {
			facVal = facVal * 10
		}
	}
	return facVal * uint32(val)
}

/*
int8:   -128 ~ 127
int16:  -32768 ~ 32767
int32:  -2147483648 ~ 2147483647
int64:  -9223372036854775808 ~ 9223372036854775807

uint8:  0 ~ 255
uint16: 0 ~ 65535
uint32: 0 ~ 4294967295
uint64: 0 ~ 18446744073709551615

dec hex 字符
48	30	0
49	31	1
50	32	2
51	33	3
52	34	4
53	35	5
54	36	6
55	37	7
56	38	8
57	39	9

例：“123456”中'1'==49
*/

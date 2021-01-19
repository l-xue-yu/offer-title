package main

import "strings"

/*
题目：将字符串中的空格替换成%20
因为golang中字符串不可修改长度，所以思路还是用新的字符串替换旧的字符串
*/
func main() {
	testStr := "We are happy"
	testStr = strings.ReplaceAll(testStr, " ", "%20") //1使用写好的函数
	println(testStr)
	testStr2 := replaceStr(testStr, " ", "%20") //2使用自己重写一遍的函数，思路和源码还有剑指offer上的一样
	println(testStr2)
}

/*
1.计算替换后的长度，创建新的字符串
2.从前往后依次替换
官方的ReplaceAll思路也是差不过的
*/
func replaceStr(changeStr, oldStr, newStr string) string {

	if oldStr == newStr || oldStr == "" {
		return changeStr
	}

	//遍历字符串，获取需要替换的字符数量，统计新字符串长度
	countStr := strings.Count(changeStr, oldStr)
	if countStr == 0 {
		return changeStr
	}
	//新字符串的长度
	newStrLen := len(changeStr) - countStr*len(oldStr) + len(newStr)*countStr
	newByte := make([]byte, newStrLen) //因为字符串不可变，所以用字节切片来存放字符串
	start := 0                         //旧字符串已经被复制的数量
	hasIndex := 0                      //新字符串已经复制的数量
	//替换每个旧的字符串
	for i := 0; i < countStr; i++ {
		index := start
		index = strings.Index(changeStr[start:], oldStr)
		if index == -1 {
			return changeStr
		}
		hasIndex += copy(newByte[:index], oldStr[:index])              //直接复制的内容
		hasIndex += copy(newByte[hasIndex:hasIndex+newStrLen], newStr) //要替换的字符串
		start = index + len(oldStr)
	}
	hasIndex += copy(newByte[hasIndex:], oldStr[:start]) //补齐后面的字符串
	return string(newByte)
}

/*

时间复杂度为O（n）的解法，搞定Offer就靠它了
我们可以先遍历一次字符串，这样就能统计出字符串中空格的总数，并可以由此计算出替换之后的字符串的总长度。
每替换一个空格，长度增加2，因此替换以后字符串的长度等于原来的长度加上2乘以空格数目。我们还是以前面的字符串"We are happy."为例，"We are happy."这个字符串的长度是14（包括结尾符号'\0'），里面有两个空格，因此替换之后字符串的长度是18。
*/

package chapter3

/*
题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则打印出1、2、3一直到最大的3位数即999。
*/

func main() {
	result := printN(2)
	for _, value := range result {
		print(value, ",")
	}
}

/*
知识点：
1.大数如何保存和表示，int,int16,int64很容易就溢出，所以一般会用字符串或数组
*/

//这种就是没有考虑大数的问题，maxNum如果超过int,int32,int64会溢出
func printN(n int) []int {

	var result []int
	var maxNum int

	//获取最大的n位数，maxNum
	for i := 0; i < n; i++ {
		maxNum = maxNum*10 + 9
	}

	//循环打印所有小于maxNum的数字
	for i := 1; i <= maxNum; i++ {
		result = append(result, i)
	}
	return result
}

/*
TODO 解法有点复杂，稍后求解
https://cloud.tencent.com/developer/article/1668723

*/

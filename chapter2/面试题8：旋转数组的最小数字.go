package main

import "fmt"

/*
题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
*/

/*
知识点：
1.二分查找法，要很熟悉
*/

func main() {
	//定义一个数组
	//arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	//BinaryFind(&arr, 0, len(arr) - 1, 30)

	//arr := []int{2, 5, 7, 15, 25, 28, 30}
	//BinaryFind(&arr, 0, len(arr) - 1, 7)
	//findVale:=binary(&arr, 0, len(arr) - 1, 7)
	//fmt.Println("main arr=",findVale)

	//旋转数组
	//arr2 := []int{3,4,5,1,2}
	//specialArr3:=[]int{1,2,3,4,5} //特殊情况数组本身有序
	specialArr3 := []int{1, 0, 1, 1, 1} //特殊情况：1,0,1,1,1
	minValue := findMin(specialArr3)
	fmt.Println("min value:", minValue)
}

//二分查找函数，基于递归的实现 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//判断 leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//要查找的数，范围应该在 leftIndex 到 middle+1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//要查找的数，范围应该在 middle+1 到 rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
	}
}

func binary(arr *[]int, leftIndex, rightIndex, findVal int) int {
	if leftIndex > rightIndex {
		return -1
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		return binary(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		return binary(arr, middle+1, rightIndex, findVal)
	} else {
		return middle
	}

}

//二分查找函数，基于循环的实现
func BinarySearchForLoop(arr *[9]int, val int) int {
	leftIndex := 0
	rightIndex := len(arr) - 1
	for leftIndex <= rightIndex {
		midIndex := (leftIndex + rightIndex) / 2
		if (*arr)[midIndex] > val { //中间值大于要查找的值则要查找的值的范围在 leftIndex (midIndex - 1)
			rightIndex = midIndex - 1
		} else if (*arr)[midIndex] < val { //中间值小于要查找的值则要查找的值的范围在 (midIndex + 1) rightIndex
			leftIndex = midIndex + 1
		} else {
			return midIndex
		}
	}
	return -1
}

/*
方法1：直接遍历整个数组，找到最小的
缺点：没有使用题目的数组的特性，是最笨的方法。
*/
func findMin1(data []int) int {
	if len(data) == 0 {
		return 0
	}
	minData := data[0]
	for _, value := range data {
		if minData > value {
			minData = value
		}
	}
	return minData
}

/*
使用二分法的思路：
1.设置两个指针index1,index2,分别指向slice[0],slice[len-1]
2.获取index1和index2的中间值，当中间值大于index1,则说明从index1到中间值mid都在旋转的数组中大的那边，最小值应该在mid后面，将index1移动到mid，并继续获取中间值
3.如果中间值小于index1,则说明mid在反转数组中较小那边，最小值应该在mid前面，将index2移动到mid，并继续获取中间值
4.如果index11+1=index2,那说明已经找到最小的值index2了。
特殊情况：
数组时顺序的比如1,2,3
数组只有个2个1,2或者2,1

*/
func findMin(data []int) int {
	if len(data) == 0 {
		return -1
	}
	if len(data) == 1 {
		return data[0]
	}
	if len(data) == 2 {
		if data[0] > data[1] {
			return data[1]
		} else {
			return data[0]
		}
	}
	sliceLen := len(data)

	index1 := 0
	index2 := sliceLen - 1
	mid := index1 //如果data[index1]大于date[index2]，则说明数组本身有序，比如1,2,3,4,5将index1给mid，直接返回，不用下面的循环,特殊情况1
	for data[index1] >= data[index2] {
		if index2-index1 == 1 { //正常情况下，当index1与index2相邻，则找到了最小的值就是index2
			mid = index2
			break
		}
		mid = (index2 + index1) / 2

		//特殊情况2:1,0,-1,1,2;1,1,1,0,1，只能顺序遍历查找
		if data[index1] == data[index2] && data[index1] == data[mid] {
			return findMin1(data)
		}

		if data[mid] > data[index1] {
			index1 = mid
		} else if data[mid] < data[index1] {
			index2 = mid
		}
	}
	return data[mid]
}

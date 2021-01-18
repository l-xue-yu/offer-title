package main

import "fmt"

func main() {
	var testArr [][]int

	testArr = append(testArr, []int{1, 2, 8, 9})
	testArr = append(testArr, []int{2, 4, 9, 12})
	testArr = append(testArr, []int{4, 7, 10, 13})
	testArr = append(testArr, []int{6, 8, 11, 15})

	println(Find(5, testArr))
}

/*
题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
考察点：如何在复杂的问题中找出规律，并将规律用代码实现。
结题思路：
1.根据题目描述找出规律
从右上角和左下角开始可以逐步排除，左上角和右下角不能逐步排除。

1 2 8 9
2 4 9 12
4 7 10 13
6 8 11 15
*/
func Find(searchNum int, data [][]int) bool {
	//从右上角，从右到左，逐个排除
	//防止出现异常二维数组，导致数组越界问题
	if len(data) < 0 {
		return false
	}
	if len(data[0]) < 0 {
		return false
	}

	row := len(data)        //行数
	col := len(data[0]) - 1 //列数
	result := false

	//先从最后一列开始
	for ; col >= 0; col-- {
		println("")
		//如果找到，退出全部循环
		if result {
			break
		}
		for i := 0; i < row; i++ {
			fmt.Printf("%d,%d;", data[i][col], searchNum)
			if data[i][col] == searchNum {
				result = true
			} else if data[i][col] < searchNum {
				//该列下方
				continue
			}
			//该列左侧
			break
		}
	}
	return result
}

/*
大体思路是右上角，最后一列第一行开始，以列为第一层循环，行为第二层循环，
如果testArr[i][j]==searchNum返回真，
否则如果testArr[i][j]>searchNum则说明searchNum在该列的左侧，需要移动到左边一列，
否则如果testArr[i][j]<searchNum则说明searchNum在该列的下行，需要到下一行
*/

/*
知识点：
1.二维数组怎么处理，我是使用列，行来遍历，要不然就是使用行，列来遍历。
*/

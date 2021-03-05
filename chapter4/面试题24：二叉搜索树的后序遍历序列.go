package main

import "fmt"

/*
题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。
参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
思路：
根据二叉搜索树的特性，左节点小于根节点，右节点大于根节点，对二叉搜索树进行后序遍历，则有以下规律：
后序遍历数组的最后一个元素是根节点，大于根节点的是左子树，小于根节点的是右子树。
根据此规律，只要左子树大于根节点的序列就不是二叉搜索树，只要右子树小于根节点的也不是二叉搜索树。

这个题解是最简单清晰的了。
https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/solution/gohen-rong-yi-li-jie-de-fang-fa-by-jtang/
*/

/*
知识点：
1,二叉搜索树的概念，根节点小于左子树，大于右子树
2.后序遍历，左，右，根
3.如何进行递归的查找
*/

func verifyPostOrder(postOrder []int) bool {
	//有点丧气的是，我没想清楚，切片小于1就是二叉搜索树
	if len(postOrder) <= 1 {
		return true
	}

	//确定根节点
	rootNode := postOrder[len(postOrder)-1]

	//寻找左子树
	leftLen := 0
	for i := 0; i < len(postOrder); i++ {
		if postOrder[i] > rootNode {
			leftLen = i
			break
		}
	}

	leftSlice := postOrder[:leftLen]
	rightSlice := postOrder[leftLen : len(postOrder)-1] //注意需要去掉根节点
	//判断左子树是否是二叉树
	//在获取左子树的时候就已经判断了左子树是否是二叉搜索树
	//for _,value:=range leftSlice{
	//	if value>rootNode{
	//		return false
	//	}
	//}

	//判断右子树是否是二叉搜索树
	for _, value := range rightSlice {
		if value < rootNode {
			return false
		}
	}

	return verifyPostOrder(leftSlice) && verifyPostOrder(rightSlice)
}

func main() {
	testSlice := []int{5, 7, 6, 9, 11, 10, 8}
	if verifyPostOrder(testSlice) {
		fmt.Println("是二叉搜索树")
	} else {
		fmt.Println("不是二叉搜索树")
	}
}

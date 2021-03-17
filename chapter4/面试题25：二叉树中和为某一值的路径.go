package main

import "fmt"

/*
题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。二叉树结点的定义如下：

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：前序遍历二叉树，将经过的节点值保存在栈中。当遍历到叶子节点且栈中和等于输入值则遍历栈打印值。否则递归遍历该节点的左子树和右子树，最后返回前删除栈中的栈顶元素。
*/

/*
知识点：
先序遍历
*/

func FindPath(tree *TreeNode, expectedSum int, path []int) {
	//二叉树为空
	if tree == nil {
		return
	}

	findPath(tree, expectedSum, path, 0)
}

/*

 */
func findPath(tree *TreeNode, expectedSum int, path []int, currentSum int) {

	if tree == nil {
		return
	}
	//当前路径值
	currentSum = currentSum + tree.Val
	path = append(path, tree.Val)
	//println("当前 sum:",currentSum)

	//是叶子节点
	if tree.Left == nil && tree.Left == tree.Right {
		if currentSum == expectedSum { //且累计值正好是输入值
			for _, value := range path {
				fmt.Print(value, " ")
			}

		} else {
			currentSum -= tree.Val
			path = path[:len(path)-1]
			return
		}

	}

	//不是叶子节点，继续递归遍历
	findPath(tree.Left, expectedSum, path, currentSum)
	findPath(tree.Right, expectedSum, path, currentSum)

}

func main() {
	/*
	   		10
	   	5		12
	   4	7
	*/
	rootTree := TreeNode{
		Val: 10,
	}
	thirdLeft := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	thirdRight := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	secondLeft := TreeNode{
		Val:   5,
		Left:  &thirdLeft,
		Right: &thirdRight,
	}
	secondRight := TreeNode{
		Val:   12,
		Left:  nil,
		Right: nil,
	}
	rootTree.Left = &secondLeft
	rootTree.Right = &secondRight
	path := make([]int, 0)
	FindPath(&rootTree, 22, path)

}

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
1.二叉树先序遍历
2.怎样设计递归函数，比如本题中如何将路径中无效的点删除。什么时候返回。什么时候开始递归
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
	if tree.Left == nil && tree.Left == tree.Right && currentSum == expectedSum {
		//且累计值正好是输入值
		for _, value := range path {
			fmt.Print(value, " ")
		}

	}

	//不是叶子节点，继续递归遍历
	findPath(tree.Left, expectedSum, path, currentSum)
	findPath(tree.Right, expectedSum, path, currentSum)
	//已经找到一个路径，将该路径上的节点一个，一个删除
	//两种情况会删除末尾点，1.找到一个符合条件的路径，在递归函数栈中将栈头的函数一个一个结束的时候删除函数中的点。2.是叶子节点，也走到了栈顶，需要结束函数栈，将栈顶函数结束，删除栈顶函数的点。
	currentSum -= tree.Val
	path = path[:len(path)-1]
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

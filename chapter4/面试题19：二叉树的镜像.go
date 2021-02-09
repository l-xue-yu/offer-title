package main

import "fmt"

/*
题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。二叉树结点的定义如下：
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
知识点：
1.先序遍历
*/

/*
思路：
遍历，然后交换左右子树，接着递归遍历子树并交换

*/

func mirrorTreeNode(pNode *TreeNode) *TreeNode {
	if pNode == nil {
		return nil
	}

	if pNode.Left == pNode.Right && pNode.Left == nil {
		return nil
	}

	//主要逻辑
	//当左右子树都存在，交换左右子树
	tmpNode := pNode.Left
	pNode.Left = pNode.Right
	pNode.Right = tmpNode

	//递归左子树，进行交转
	if pNode.Left != nil {
		mirrorTreeNode(pNode.Left)
	}
	//递归右子树，进行交换
	if pNode.Right != nil {
		mirrorTreeNode(pNode.Right)
	}

	return pNode
}

func preOrder(pNode *TreeNode) {
	if pNode == nil {
		return
	}
	fmt.Print(pNode.Val, " ")
	preOrder(pNode.Left)
	preOrder(pNode.Right)
}

func main() {

	ALeft2 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	ALfet3 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	Aright3 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	ARight2 := &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	}
	ARight1 := &TreeNode{
		Val:   10,
		Left:  ALfet3,
		Right: ARight2,
	}
	ALeft1 := &TreeNode{
		Val:   6,
		Left:  ALeft2,
		Right: Aright3,
	}
	A1 := &TreeNode{
		Val:   8,
		Left:  ALeft1,
		Right: ARight1,
	}

	preOrder(A1)

	mirrorA1 := mirrorTreeNode(A1)
	fmt.Println("")
	preOrder(mirrorA1)
}

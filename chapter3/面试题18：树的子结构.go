package main

/*
题目：输入两棵二叉树A和B，判断B是不是A的子结构。二叉树结点的定义如下
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
知识点：
1.如何遍历二叉树，递归遍历，循环遍历，
二叉树的遍历，特别是递归遍历，说到底是对根节点，左子树，右子树分别进行递归遍历，不过根据顺序不同，有前序（先根节点，然后左子树，最后右子树），中序（先左子树，然后根节点，最后右子树），后序（先左子树，然后右子树，最后根节点）区分。
也就是说，根据二叉树的结构特性，二叉树的递归遍历一定是对根节点，左子树，右子树进行遍历。

*/

/*
思路：
1.isSubStructure函数来递归遍历二叉树A，在遍历中，如果节点的值相同，对A,B都进行遍历，判断是否左右子树的各个值是否相同，
如果节点值不同，则继续递归遍历二叉树A并传入B，同时在后续遍历中，重复上面的步骤，即找到相同的值，进行深层遍历，找不到进行对A进行递归遍历
2.深层遍历，判断A的子树是否和B相同，相同进行下一个深层遍历，不相同返回false，如果b==nil说明已经找完B所有的子节点且都相同，返回true，如果a==nil返回失败

*/

func isSubStruct(A, B *TreeNode) bool {

	//两个都到了原子节点，说明该节点已经遍历完成
	if A == nil && B == nil {
		return true
	}

	//一个为空，说明，该节点不完全相同，返回false
	if A == nil || B == nil {
		return false
	}

	res := false
	//节点值相同，进行深层递归遍历
	if A.Val == B.Val {
		res = deepFind(A, B)
	}

	//res==false，说明上层节点不相同，比较节点的下一层的左子树
	if !res {
		res = isSubStruct(A.Left, B)
	}
	//res==false，说明上层节点不相同，比较节点的下一层的右子树子树
	if !res {
		res = isSubStruct(A.Right, B)
	}

	return res
}

func deepFind(a, b *TreeNode) bool {

	//前面所有的递归都是true，那么意思现在是找到最后一个原子节点了，
	if b == nil {
		return true
	}

	//b!=nil,而A==nil,意思不相等，直接返回false
	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}
	//递归校验A B左子树和右子树的结构和节点是否相同
	return deepFind(a.Left, b.Left) && deepFind(a.Right, b.Right)

}

/*
基本是参考的力扣的题解
https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/solution/golang-di-gui-by-sakura-151/
*/

func main() {

	//二叉树A
	ARight1 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	ALeft2 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	ALfet3 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	Aright3 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	ARight2 := &TreeNode{
		Val:   2,
		Left:  ALfet3,
		Right: Aright3,
	}
	ALeft1 := &TreeNode{
		Val:   8,
		Left:  ALeft2,
		Right: ARight2,
	}
	A1 := &TreeNode{
		Val:   8,
		Left:  ALeft1,
		Right: ARight1,
	}

	//二叉树B
	BLeft2 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	BRight2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	B1 := &TreeNode{
		Val:   8,
		Left:  BLeft2,
		Right: BRight2,
	}

	//二叉树C
	CLeft2 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	CRight2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	C1 := &TreeNode{
		Val:   8,
		Left:  CLeft2,
		Right: CRight2,
	}

	if isSubStruct(A1, B1) {
		println("B节点是A节点的子树")
	} else {
		println("B节点不是A节点的子树")
	}
	if isSubStruct(A1, C1) {
		println("C节点是A节点的子树")
	} else {
		println("C节点不是A节点的子树")
	}

}

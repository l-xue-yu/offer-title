package main

import "fmt"

/*
题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出图2.6所示的二叉树并输出它的头结点。
二叉树结点的定义如下：
*/

//二叉树结构
type BinaryTreeNode struct {
	value   int             //数据
	mpLeft  *BinaryTreeNode //左子树
	mpRight *BinaryTreeNode //右子树
}

func main() {
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	newTreeNode := buildTreeNode(preorder, inorder)
	newTreeNode.PreOrder()
	newTreeNode.middleOrder()

}

func buildTreeNode(preroder, inorder []int) *BinaryTreeNode {

	//先序切片用完或中序用完退出
	if len(preroder) == 0 || len(inorder) == 0 {
		//fmt.Println("到底原子节点")
		return nil
	}
	//根据前序找根节点
	rootValue := preroder[0]

	//初始化二叉树
	newTreeNode := new(BinaryTreeNode)
	newTreeNode.value = rootValue
	newTreeNode.mpLeft = nil
	newTreeNode.mpRight = nil

	//遍历中序，获取先序遍历的左子树和右子树，中序遍历的左子树和右子树

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootValue {
			//每次使用拷贝的底层数组的切片而不是本身的切片
			tmpPreorder := make([]int, len(preroder))
			tmpInorder := make([]int, len(preroder))
			tmpPreIndex := copy(tmpPreorder, preroder)
			tmpIndex := copy(tmpInorder, inorder)
			if tmpPreIndex == 0 || tmpIndex == 0 {
				return nil
			}
			//递归调用自身，传入先序和中序，补全下面的二叉树
			//注：确定左子树，右子树的公式，最好在纸上画出来，详情见剑指offer面试题6
			newTreeNode.mpLeft = buildTreeNode(tmpPreorder[1:i+1], tmpInorder[:i])   //左子树的先序 tmpPreorder[1:i+1]，左子树的中序 tmpInorder[:i]
			newTreeNode.mpRight = buildTreeNode(tmpPreorder[i+1:], tmpInorder[i+1:]) //右子树的先序 tmpPreorder[i+1:]，左子树的中序tmpInorder[i+1:]
		}
	}

	return newTreeNode
}

//打印节点
func (node *BinaryTreeNode) Print() {
	fmt.Print(node.value, " ")
}

//前序遍历： 根 前序左  前序右
func (node *BinaryTreeNode) PreOrder() {
	if node == nil {
		return
	}

	node.Print()
	node.mpLeft.PreOrder()
	node.mpRight.PreOrder()
}

//中序遍历：左 中 右
func (node *BinaryTreeNode) middleOrder() {
	if node == nil {
		return
	}

	node.mpLeft.middleOrder()
	node.Print()
	node.mpRight.middleOrder()
}

/*
前序
1 2 4 7 3 5 6 8
中序
4 7 2 1 5 3 8 6

逻辑过程：
1.根据前序，确定1是根节点
2.遍历中序，确定左子树的范围是4 7 2,即长度是3
3.那么可以获取到左子树的前序2 4 7 ，中序是4 7 2
5.递归调用自己，将所有左子树的节点补充
6.将右子树的前序3 5 6 8,中序5 3 8 6获取到，然后递归调用，补充所有右子树节点
7.左子树补充完，右子树补充完后，返回二叉树

*/

/*
知识点：
1.二叉树结构
2.二叉树的，前序遍历，中序遍历，后序遍历，要非常熟悉
3.根据前序遍历，确定根节点，根据中序遍历，确定左子树和右子树范围，并确定左子树的前序遍历，左子树的中序遍历，
4.用递归的方法来完成二叉树的构建，这里所说的构建指通过递归将所有节点的数据域，左子树，右子树全部赋值。
5.切片复制，使用copy来复制底层数组，而不是复制切片本身的结构
*/

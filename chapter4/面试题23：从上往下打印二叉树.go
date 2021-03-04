package main

import "fmt"

/*
题目：从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。例如输入图4.5中的二叉树，则依次打印出8、6、10、5、7、9、11。
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7


 Definition for a binary tree node.


*/

/*
思路：利用队列的特性，先入先出，从根节点开始，然后左右节点，节点不为空将节点的下层左右节点入队。
如：入队3,9,20，出队3，打印9,9没有下层节点直接出队，打印20,20有下层节点，将15,7入队，20出队，打印15,15没有下层节点，直接出队。
步骤		操作				队列
1			入队3,9,20			3,9,20
2			打印3，出队3			9,20
3			打印9,出队9			20
4			打印20，入队15,7		15,7
5			打印15，出队15		7
6			打印7出队7
*/

/*
知识点：
1.二叉树和队列的特性。
2.使用辅助的容器（当正常的结构不能完成时，添加新的容器）。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
有两种思路，出队和不出队都行，出队的话就要将下一个节点给当前节点，不出队就直接找队列的下一个
*/
func levelOrder(root *TreeNode) (res []int) {

	//异常数据返回
	if root == nil {
		return res
	}
	//根节点入队
	q := []*TreeNode{root}

	//当前节点
	var presentNode *TreeNode

	for i := 0; i < len(q); i++ {
		presentNode = q[i]
		fmt.Println(presentNode.Val) //打印节点
		res = append(res, presentNode.Val)
		if presentNode.Left != nil { //判断左子树是否存在，存在入队
			q = append(q, presentNode.Left)
		}

		if presentNode.Right != nil { //判断右子树是否存在，存在入队
			q = append(q, presentNode.Right)
		}

	}
	return res
}

func main() {

	ALfet3 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	ARight2 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	ARight1 := &TreeNode{
		Val:   20,
		Left:  ALfet3,
		Right: ARight2,
	}
	ALeft1 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	A1 := &TreeNode{
		Val:   3,
		Left:  ALeft1,
		Right: ARight1,
	}

	_ = levelOrder(A1)
}

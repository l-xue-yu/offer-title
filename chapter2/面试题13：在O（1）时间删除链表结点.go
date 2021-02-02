package main

import "fmt"

/*
面试题13：在O（1）时间删除链表结点
*/

/*
知识点：
1.删除节点的两种方式，遍历删除，时间复杂度使O(n)，节点覆盖，将后一个节点的内容覆盖到要删除的节点，不需要遍历节点。
*/

//链表节点定义
type ListNode struct {
	mnValue int
	mpNext  *ListNode
}

func main() {
	newListNode1 := &ListNode{
		mnValue: 1,
		mpNext:  nil,
	}
	newListNode2 := &ListNode{
		mnValue: 2,
		mpNext:  nil,
	}
	newListNode3 := &ListNode{
		mnValue: 3,
		mpNext:  nil,
	}
	newListNode1.mpNext = newListNode2
	newListNode2.mpNext = newListNode3
	printListNode(newListNode1)

	DeleteNode(newListNode1, newListNode1)

	printListNode(newListNode1)
}

func DeleteNode(headNode *ListNode, toDeleteNode *ListNode) {
	//删除的如果是头结点该链表只有一个节点
	if headNode == toDeleteNode && toDeleteNode.mpNext == nil {
		headNode = nil
		toDeleteNode = nil
		return
	}

	//删除的如果是尾结点,还是要遍历
	if toDeleteNode.mpNext == nil {
		pNode := headNode
		for pNode.mpNext != nil {
			if pNode.mpNext == toDeleteNode {
				pNode.mpNext = nil
				return
			}
			pNode = pNode.mpNext
		}
	}

	//删除的是中间的普通节点
	//将下一个节点的数据复制到该节点
	toDeleteNode.mnValue = toDeleteNode.mpNext.mnValue
	toDeleteNode.mpNext = toDeleteNode.mpNext.mpNext
}

//遍历打印节点
func printListNode(headNode *ListNode) {
	pNode := headNode
	for pNode != nil {
		fmt.Println(pNode.mnValue)
		pNode = pNode.mpNext
	}

}

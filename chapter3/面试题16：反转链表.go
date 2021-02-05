package main

import "fmt"

/*
题目：定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的头结点。链表结点定义如下：
*/

/*
之前已经整理过链表反转,画了张图，比较简单易懂，地址：https://blog.csdn.net/qq_26372385/article/details/103567221
*/

type ListNode4 struct {
	mnValue int
	mpNext  *ListNode3
}

//单链表反转1，循环反转
//链表需要节点的相关节点如下：
//前驱节点，predecessorNode
//当前节点，currentNode
//后续节点，nextCode
func ReverseCycle(head *ListNode3) {
	if head == nil || head.mpNext == nil {
		return
	}

	var predecessorNode, nextNode *ListNode3 //前驱节点和后续节点都应该为空节点（nil）
	currentNode := head                      //当前节点
	for currentNode != nil {
		nextNode = currentNode.mpNext        //1保存后续节点，防止节点丢失
		currentNode.mpNext = predecessorNode //2断链并将当前节点指向前驱节点，完成该节点的反转

		predecessorNode = currentNode //3将当前节点给前驱节点
		currentNode = nextNode        //4将原后续节点给当前节点，完成节点的易懂
	}
}

func main() {
	newNode1 := &ListNode3{
		mnValue: 1,
		mpNext:  nil,
	}
	newNode2 := &ListNode3{
		mnValue: 2,
		mpNext:  nil,
	}
	newNode3 := &ListNode3{
		mnValue: 3,
		mpNext:  nil,
	}
	newNode4 := &ListNode3{
		mnValue: 4,
		mpNext:  nil,
	}
	newNode5 := &ListNode3{
		mnValue: 5,
		mpNext:  nil,
	}
	newNode1.mpNext = newNode2
	newNode2.mpNext = newNode3
	newNode3.mpNext = newNode4
	newNode4.mpNext = newNode5

	headNode := newNode1
	//for headNode!=nil{
	//	fmt.Printf("%v",headNode.mnValue)
	//	headNode=headNode.mpNext
	//}
	headNode = newNode1
	ReverseCycle(headNode)
	//println("反转后：")
	headNode = newNode5
	for headNode != nil {
		fmt.Printf("%v", headNode.mnValue)
		headNode = headNode.mpNext
	}
}

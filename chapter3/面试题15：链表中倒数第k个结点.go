package main

import "fmt"

/*
题目：输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是值为4的结点。
*/

type ListNode2 struct {
	mnValue int
	mpNext  *ListNode2
}

/*
思路：
定义两个指针，第一个指针先走，第二个指针等第一个指针走了k-1步才出发，当第一个指针走到尾部，第二个指针就走到了倒数第k个节点
*/

/*
知识点：
1.强调鲁棒性，养成防御性编程的习惯，对可能导致程序异常，错误或崩溃的输入提前做验证，返回提示
*/

func FindKthToTail(node *ListNode2, k int) *ListNode2 {

	//预留异常输入验证
	if node == nil || k == 0 {
		return nil
	}

	firstNode := node
	secNode := node
	step := 0
	for firstNode.mpNext != nil {
		if k-1 <= step {
			secNode = secNode.mpNext
		}
		firstNode = firstNode.mpNext
		step++
	}
	return secNode
}

func main() {
	newNode1 := &ListNode2{
		mnValue: 1,
		mpNext:  nil,
	}
	newNode2 := &ListNode2{
		mnValue: 2,
		mpNext:  nil,
	}
	newNode3 := &ListNode2{
		mnValue: 3,
		mpNext:  nil,
	}
	newNode4 := &ListNode2{
		mnValue: 4,
		mpNext:  nil,
	}
	newNode5 := &ListNode2{
		mnValue: 5,
		mpNext:  nil,
	}
	newNode1.mpNext = newNode2
	newNode2.mpNext = newNode3
	newNode3.mpNext = newNode4
	newNode4.mpNext = newNode5

	resNode := FindKthToTail(newNode1, 3)
	fmt.Printf("%v", resNode)
}

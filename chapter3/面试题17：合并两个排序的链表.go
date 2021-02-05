package main

import "fmt"

/*
题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的。例如输入图3.7中的链表1和链表2，则合并之后的升序链
*/

/*
思路：
如
P1 1,3,5,7
P2 2,4,6,8
P3 1,2,4,5,6,7,8
如果P1node<P2Node{
	p3Node=p1Node
	p3Node.next=merge(p1Node.next,p2Node)

}else{
	p3Node=p2Node
	p3Node.next=merge(p1Node.next,p2Node.next)
}


*/

/*
知识点1：如何将整个流程抽象成一个递归函数，这道题思路很简单，就是输入两个单链表，链表本身已经排序，只要把链表逐个比对，小的给合并的链表，剩下的进行下一次比对。
知识点2：如何写递归函数，1将函数的主要实现步骤，抽象出来，2，递归函数的输入，3递归函数的返回，递归函数什么时候结束
*/

type ListNode4 struct {
	mnValue int
	mpNext  *ListNode4
}

func merge(head1, head2 *ListNode4) *ListNode4 {

	//预留异常数据验证
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	var headMerge *ListNode4 //新建一个节点

	if head1.mnValue < head2.mnValue {
		headMerge = head1                             //将head1给headMerge
		headMerge.mpNext = merge(head1.mpNext, head2) //进行下一个比较，因为head1已经比较过了，将head1.next和head进行比较
	} else {
		headMerge = head2
		headMerge.mpNext = merge(head1, head2.mpNext)
	}
	return headMerge
}

func main() {
	newNode1 := &ListNode4{
		mnValue: 1,
		mpNext:  nil,
	}
	newNode2 := &ListNode4{
		mnValue: 3,
		mpNext:  nil,
	}
	newNode3 := &ListNode4{
		mnValue: 5,
		mpNext:  nil,
	}
	newNode4 := &ListNode4{
		mnValue: 7,
		mpNext:  nil,
	}
	newNode5 := &ListNode4{
		mnValue: 9,
		mpNext:  nil,
	}
	newNode1.mpNext = newNode2
	newNode2.mpNext = newNode3
	newNode3.mpNext = newNode4
	newNode4.mpNext = newNode5

	newNodeB1 := &ListNode4{
		mnValue: 2,
		mpNext:  nil,
	}
	newNodeB2 := &ListNode4{
		mnValue: 4,
		mpNext:  nil,
	}
	newNodeB3 := &ListNode4{
		mnValue: 6,
		mpNext:  nil,
	}
	newNodeB4 := &ListNode4{
		mnValue: 8,
		mpNext:  nil,
	}
	newNodeB5 := &ListNode4{
		mnValue: 10,
		mpNext:  nil,
	}
	newNodeB1.mpNext = newNodeB2
	newNodeB2.mpNext = newNodeB3
	newNodeB3.mpNext = newNodeB4
	newNodeB4.mpNext = newNodeB5

	nodeC1 := merge(newNode1, newNodeB1)
	for nodeC1 != nil {
		fmt.Printf("%v", nodeC1.mnValue)
		nodeC1 = nodeC1.mpNext
	}
}

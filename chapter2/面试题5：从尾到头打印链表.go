package main

import "fmt"

/*
题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。链表结点定义如下：

*/

type listNode struct {
	data int
	next *listNode
}

func main() {
	//链表
	n1 := new(listNode)
	n1.data = 1
	n2 := new(listNode)
	n2.data = 2
	n1.next = n2
	n3 := new(listNode)
	n3.data = 3
	n2.next = n3
	n3.next = nil

	//printList(n1)
	printCycle(n1)

}

/*
循环实现
思路：遍历链表，然后把从头到尾的数据入栈，最后打印栈
*/
func printList(headNode *listNode) {
	sliceStack := make([]int, 0)
	for headNode != nil {
		//入栈
		sliceStack = append(sliceStack, headNode.data)
		headNode = headNode.next
	}

	//打印栈
	for i := len(sliceStack) - 1; i >= 0; i-- {
		fmt.Print(sliceStack[i])
	}

}

/*
递归实现

*/
func printCycle(list *listNode) {
	if list != nil {
		if list.next != nil {
			printCycle(list.next)
		}
		fmt.Printf("%d", list.data)
	}
}

//用切片来实现栈
/*
入栈：slice=append(slice,data)
遍历栈：
for i:=slice.length-1;i>=0;i--{
	fmt.print(slice[i])
}
*/

/*
知识点：
1.链表，链表的结构，指针域，数据域，如何遍历链表
2.栈，栈的结构，如何入栈，出栈，遍历栈
3.递归,用递归来遍历链表
*/

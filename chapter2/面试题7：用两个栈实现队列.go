package main

import "fmt"

/*
题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail和deleteHead，分别完成在队列尾部插入结点和在队列头部删除结点的功能。
*/

/*
知识点：
1.队列的概念，先进先出
2.队列操作，入队尾部插入，出队头部删除
3.可以使用切片来实现队列的操作
*/

/*
思路：
入队1 2 3 4 5后
in   1 2 3 4 5
out

出队
in
out 5 4 3 2 1
出队后
in
out 5 4 3 2

1.in作为入队的栈,入队就直接append到栈尾
2.out作为出队的栈，要想实现出队的先入先出，就是将in栈的元素倒叙压入out栈，然后出队从队尾出，就是先入先入

问题：这样写应该会有竞态问题，特别是出队，最好是加锁。
*/

//定义队列的结构体
type queue struct {
	in  []int //入队的栈
	out []int //出队的栈
}

func main() {
	newQueue := new(queue)

	for i := 1; i < 6; i++ {

		newQueue.appendTail(i)

		fmt.Printf("入队栈：%v", newQueue.in)
	}
	println("")
	for i := 1; i < 6; i++ {
		newQueue.deleteHead()
		fmt.Printf("出队栈：%v", newQueue.out)
	}
}

//入队
func (thisQueue *queue) appendTail(data int) {
	thisQueue.in = append(thisQueue.in, data)
}

//出队,出队栈没有元素，先查询入队栈是否有元素，有元素先倒叙压入，没有直接返回。出队栈有元素后，执行出栈，每次将栈最后一个元素删除，直到没有元素了，再从入队栈压入元素。
func (thisQueue *queue) deleteHead() int {

	if len(thisQueue.out) == 0 {
		//没有元素，需要从入队的栈压入
		if len(thisQueue.in) == 0 {
			return -1
		}

		for i := len(thisQueue.in) - 1; i >= 0; i-- {
			thisQueue.out = append(thisQueue.out, thisQueue.in[i])
		}
	}
	thisQueue.out = thisQueue.out[:len(thisQueue.out)-1]
	return 0
}

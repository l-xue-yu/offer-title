package main

/*
题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min函数。在该栈中，调用min、push及pop的时间复杂度都是O（1）。
*/

/*
思路：
还是找规律，感觉并不是很容易就能想出方法。
结题思路：只要能保证压入辅助栈的元素是小于等于上一个元素，那么辅助栈就是有序的，每次弹出就是最小的元素。
4			4		4
43			43		3
435			433		3
疑惑，切片不是并发安全的，按照正常的思路写，感觉会有隐患……
*/

/*
知识点：
1.栈用切片来作为底层结构
*/

type MinStack struct {
	nums []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		nums: make([]int, 0, 3),
		min:  make([]int, 0, 3),
	}
}

//入栈
func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)

	if len(this.min) == 0 { //辅助栈没有元素直接入栈
		this.min = append(this.min, x)
	} else if x < this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

//出栈
func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Min() int {
	return this.min[this.min[len(this.min)-1]]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

package main

import "fmt"

/*
题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
*/

/*
第一种思路：
从头扫描这个数组，每碰到一个偶数时，拿出这个数字，并把位于这个数字后面的所有数字往前挪动一位。挪完之后在数组的末尾有一个空位，这时把该偶数放入这个空位。由于每碰到一个偶数就需要移动O（n）个数字，因此总的时间复杂度是O（n2）。
*/

/*
第二种思路：
维护两个指针，第一个指针初始化时指向数组的第一个数字，它只向后移动；第二个指针初始化时指向数组的最后一个数字，它只向前移动。在两个指针相遇之前，第一个指针总是位于第二个指针的前面。如果第一个指针指向的数字是偶数，并且第二个指针指向的数字是奇数，我们就交换这两个数字。
以一个具体的例子比如输入数组{1,2,3,4,5}来分析这种思路。在初始化时，把第一个指针指向数组第一个数字1，而把第二个指针指向最后一个数字5（如图3.4（a）所示）。第一个指针指向的数字1是一个奇数，不需要处理，我们把第一个指针向后移动，直到碰到一个偶数2。
此时第二个指针已经指向了奇数，因此不需要移动。此时两个指针指向的位置如图3.4（b）所示。这个时候我们发现偶数2位于奇数5的前面，符合我们的交换条件，于是交换两个指针指向的数字
*/

/*
感觉这种思路和快排的哨兵很像，两个指针，分别从头，尾，往中间走，当两个指针都满足条件开始交换,直到相遇或交替，结束
*/

func changeOrder(data []int) {

	left := 0
	right := len(data) - 1

	for left < right {
		if data[left]&1 == 1 { //奇数
			left++
		} else if data[right]&1 == 1 {
			//left 是偶数，right是奇数,交换
			tmp := data[left]
			data[left] = data[right]
			data[right] = tmp
		} else {
			right--
		}
	}

}

func main() {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("before order:%v", testSlice)
	changeOrder2(testSlice)
	fmt.Printf("after order:%v", testSlice)
}

/*
可扩展性优化，解耦，将操作分成两步，第一步判断是否满足标准，第二步交换
下面的写法是推荐的写法，但是我感觉这里left<right的判断太多了，我上面写的全if逻辑复杂了，而且for的次数会比下面的多。
*/

func changeOrder2(data []int) {

	left := 0
	right := len(data) - 1

	for left < right {

		for left < right && !isEven(data[left]) {
			//奇数
			left++
		}
		for left < right && isEven(data[right]) {
			//偶数
			right--
		}

		if left < right {
			//left 是偶数，right是奇数,交换
			tmp := data[left]
			data[left] = data[right]
			data[right] = tmp
		}

	}

}

func isEven(data int) bool {
	//偶数返回true
	return data&1 == 0
}

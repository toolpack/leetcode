package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(85)
	obj.Push(-99)
	obj.Push(67)

	fmt.Println(obj.GetMin())
	// obj.Pop()
	// fmt.Println(obj.Top())
	// fmt.Println(obj.GetMin())
}

type MinStack struct {
	s1 []int
	s2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s1 = append(this.s1, x)
	if len(this.s2) == 0 {
		this.s2 = append(this.s2, x)
	} else if x <= this.s2[len(this.s2)-1] {
		this.s2 = append(this.s2, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.s1) == 0 {
		return
	}
	data := this.s1[len(this.s1)-1]
	this.s1 = this.s1[:len(this.s1)-1]
	if data == this.GetMin() {
		this.s2 = this.s2[:len(this.s2)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.s1) == 0 {
		return 0
	}
	return this.s1[len(this.s1)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.s2) == 0 {
		return 0
	}
	return this.s2[len(this.s2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

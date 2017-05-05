package main

func main() {

}

type MyQueue struct {
	l []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := MyQueue{}
	return q
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.l = append(this.l, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.l) == 0 {
		return 0
	}
	x := this.l[0]
	this.l = this.l[1:]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.l) == 0 {
		return 0
	}
	x := this.l[0]
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.l) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

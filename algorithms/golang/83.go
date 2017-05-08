package main

import (
	"fmt"
)

func main() {
	l2 := &ListNode{1, nil}
	l1 := &ListNode{1, l2}
	// l3 := l1
	// fmt.Printf("%p val:%v\n", &l1, l1.Val)
	// fmt.Printf("%p val:%v\n", &l3, l3.Val)

	//fmt.Printf("%p val:%p\n", &l1, &l1.Val)
	//fmt.Printf("%p val:%p\n", &l3, &l3.Val)

	deleteDuplicates(l1)
	l1.print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Printf("value:%v, p: %p\n", l.Val, &l)
		l = l.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[int]bool)
	m[head.Val] = true
	pre := head
	p := head.Next
	for p != nil {
		if _, ok := m[p.Val]; !ok {
			m[p.Val] = true
			pre = p
		} else {
			pre.Next = p.Next
		}
		p = p.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {

	return head
}

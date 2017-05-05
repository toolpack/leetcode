package main

import (
	"fmt"
	"time"
)

func main() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	// n1.print()
	l := reverselink(n1)
	l.print()
	// n1.print()
	// fmt.Println(n1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for l != nil {
		time.Sleep(time.Second)
		fmt.Println(l.Val)
		l = l.Next
	}
}

/*
func isPalindrome(head *ListNode) bool {
	l := []int{}
	for head != nil {
		l = append(head.Val)
		head = head.Next
	}
	for i := 0; i < len(l)/2; i++ {
		if l[i] != l[len(l)-i-1] {
			return false
		}
	}
	return true
}*/
/*
func reverselink(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	p := head
	q := head.Next
	var r *ListNode
	head.Next = nil
	for q != nil {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}
	// p.print()
	// *head = *p
}

//1 2 3 4
//
*/

func reverselink(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	p := h
	q := h.Next
	h.Next = nil
	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}
	return p
}

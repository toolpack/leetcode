package main

import (
	"fmt"
)

func main() {
	n5 := &node{5, nil}
	n4 := &node{4, n5}
	n3 := &node{3, n4}
	n2 := &node{2, n3}
	n1 := &node{1, n2}
	do328(n1)
	fmt.Println(n1.data)
	fmt.Println(n1.next.data)
	fmt.Println(n1.next.next.data)
	fmt.Println(n1.next.next.next.data)
	fmt.Println(n1.next.next.next.next.data)
}

type node struct {
	data int
	next *node
}

func do328(n *node) {
	if n == nil || n.next == nil {
		return
	}
	odd := n
	even := n.next
	evenHead := even

	for odd != nil && odd.next != nil {
		odd.next = odd.next.next
		even.next = even.next.next
		odd = odd.next
		even = even.next
	}
	odd.next = evenHead
}

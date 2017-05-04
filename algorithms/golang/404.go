package main

func main() {
	n1 := node{15, nil, nil}
	n2 := node{7, nil, nil}
	n3 := node{20, &n1, &n2}
	n4 := node{9, nil, nil}
	n5 := node{3, &n4, &n3}
	tree(&n5)
	print(sum)
}

type node struct {
	data int
	l    *node
	r    *node
}

var sum int

func tree(n *node) {
	if n == nil {
		return
	}
	if n.l != nil {
		sum += n.l.data
		tree(n.l)
	}
	if n.r != nil {
		tree(n.r)
	}
}

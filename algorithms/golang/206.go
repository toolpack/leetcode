package main

import (
	"fmt"
)

func main() {
	l := []int{4, 3, 2, 1}
	qsort(l, 0, len(l)-1)
	fmt.Println(l)
}

type node struct{}

func do206(n *node) *node {
	return nil
}

func maopao(l []int) []int {
	if len(l) == 0 {
		return l
	}
	for i := 0; i < len(l)-1; i++ {
		for j := 0; j < len(l)-1-i; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
	return l
}

func qsort(l []int, left int, right int) {
	if left < right {
		m := l[left]
		i, j := left, right
		for {
			for l[j] >= m && i < j {
				j--
			}
			for l[i] <= m && i < j {
				i++
			}
			if i >= j {
				break
			}
			l[i], l[j] = l[j], l[i]
		}
		l[left] = l[i]
		l[i] = m
		qsort(l, left, i-1)
		qsort(l, i+1, right)
	}
}

package main

func main() {
	do303([]int{-2, 0, 3, -5, 2, -1}, 0, 2)
	do303([]int{-2, 0, 3, -5, 2, -1}, 2, 5)
	do303([]int{-2, 0, 3, -5, 2, -1}, 0, 5)
}

func do303(l []int, i, j int) {
	sum := 0
	for ; i <= j; i++ {
		sum += l[i]
	}
	print(sum)
	print("\n")
}

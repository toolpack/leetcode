package golang

func plusOne(n []int) []int {
	r := []int{}

	carry := 1
	for i := len(n) - 1; i >= 0; i-- {
		x := n[i] + carry
		y := x % 10
		carry = x / 10
		r = append([]int{y}, r...)
	}
	if carry != 0 {
		r = append([]int{carry}, r...)
	}
	return r
}

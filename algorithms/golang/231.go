package main

func main() {

}
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for ; n > 1; n /= 2 {
		x := n % 2
		if x != 0 {
			return false
		}
	}
	return true
}

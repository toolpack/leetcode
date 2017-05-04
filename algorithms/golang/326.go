package main

func main() {
	print(do326(28))
	print(do326(256))
}

func do326(n int) bool {
	for ; n > 1; n /= 3 {
		x := n % 3
		if x != 0 {
			return false
		}
	}
	return true
}

package main

func main() {
	do263(6)
	do263(8)
	do263(14)
}

func do263(n int) {
	for i := 2; i < 6 && n > 0; i++ {
		for {
			if n%i == 0 {
				n /= i
			} else {
				break
			}
		}
	}
	print(n == 1)
}

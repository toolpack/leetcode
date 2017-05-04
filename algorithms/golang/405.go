package main

func main() {
	tohex(26)
}

func tohex(i int) string {
	if i == 0 {
		return "0"
	}
	var result string
	for ; i > 0; i /= 16 {
		n := byte(i % 16)
		var c byte
		if n < 10 {
			c = n + '0'
		} else {
			c = n - 10 + 'a'
		}
		result = string(c) + result
	}
	print(result)
	return result
}

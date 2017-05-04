package main

func main() {
	sum("1111111111111111111111111111111111", "2222222222222222222222222222222222222222")
}

func sum(a, b string) string {
	var short, long string

	if len(a) > len(b) {
		long = a
		short = b
	} else {
		long = b
		short = a
	}

	var carry byte
	var result string
	i, j := len(short)-1, len(long)-1
	for j > 0 {
		var add byte
		i--
		j--
		if i >= 0 {
			add = short[i] + long[i] - 2*'0' + carry
		} else {
			add = long[j] - '0' + carry
		}
		result = string(add%10+'0') + result
		carry = add / 10
	}
	print(result)
	return result
}

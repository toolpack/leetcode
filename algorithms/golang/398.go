package main

var l = []int{1, 2, 3, 3, 3}

func main() {
	pick(3)
}

func pick(n int) {
	var tmp []int
	for _, v := range l {
		if v == n {
			tmp = append(tmp, v)
		}
	}
	//random pick
}

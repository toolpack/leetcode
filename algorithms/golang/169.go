package main

func main() {

}
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	max := 0
	r := 0
	for k, v := range m {
		if v > max {
			max = v
			r = k
		}
	}
	return r
}

func majorityElement2(nums []int) int {
	m := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			m = nums[i]
			cnt++
		} else {
			if m == nums[i] {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return 0
}

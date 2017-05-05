package main

func main() {

}
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	m := make(map[int]int)
	for k0, v := range nums {
		tmp, ok := m[v]
		if ok && k0-tmp <= k {
			return true
		}
		m[v] = k0
	}
	return false
}

package golang

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := len(nums)
	m := nums[0]
	cnt := 1
	for i := 1; i < l; i++ {
		if nums[i] != m {
			cnt++
			m = nums[i]
		} else {
			nums = append(nums[i:], nums[i+1:]...)
			l--
			i--
		}
	}
	return cnt
}

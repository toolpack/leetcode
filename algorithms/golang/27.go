package golang

func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			l--
			i--
		}
	}
	return len(nums)
}

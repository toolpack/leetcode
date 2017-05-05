package main

import (
	"fmt"
)

func main() {
	r := summaryRanges([]int{0, 1, 2, 4, 5, 7, 9})
	fmt.Println(r)
}
func summaryRanges(nums []int) []string {
	var result []string
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		result = append(result, fmt.Sprintf("%v", nums[0]))
		return result
	}

	var begin, end, pre int
	for k, v := range nums {
		if k == 0 {
			begin = v
			pre = v
			continue
		}
		if v-1 > pre {
			end = pre
			if begin == end {
				result = append(result, fmt.Sprintf("%v", begin))
			} else {
				result = append(result, fmt.Sprintf("%v->%v", begin, end))
			}
			begin = v
			if k == len(nums)-1 {
				result = append(result, fmt.Sprintf("%v", v))
			}
		}
		pre = v
	}
	return result
}

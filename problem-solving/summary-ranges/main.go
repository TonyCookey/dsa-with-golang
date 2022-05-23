package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	low := nums[0]
	start := true
	var result []string

	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1] > nums[i]+1 {
			rang := fmt.Sprintf("%v", low)
			if !start {
				rang += "->" + fmt.Sprintf("%v", nums[i])
			}
			result = append(result, rang)
			low = nums[i+1]
			start = true
		} else {
			start = false
		}
	}
	if low == nums[len(nums)-1] {
		result = append(result, fmt.Sprintf("%v", nums[len(nums)-1]))
	} else {
		result = append(result, fmt.Sprintf("%v%s%v", low, "->", nums[len(nums)-1]))
	}
	return result
}

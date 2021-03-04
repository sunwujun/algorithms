package n0018

import (
	"sort"
)

// [18] 四数之和
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	l := len(nums)
	if l < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < l-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// 固定值去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, l-1
			for left < right {
				temp := nums[i] + nums[j] + nums[left] + nums[right]
				if temp == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}

				if temp < target {
					left++
				}

				if temp > target {
					right--
				}
			}
		}
	}
	return result
}

package n0015

import (
	"sort"
)

// [15] 三数之和
func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	var left, right, temp int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		// 固定值去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right = i+1, len(nums)-1
		for left < right {
			temp = nums[i] + nums[left] + nums[right]
			switch {
			case temp == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 指针值去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			case temp > 0:
				right--
			case temp < 0:
				left++
			}
		}
	}
	return result
}

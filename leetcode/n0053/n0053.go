package n0053

// 动态规划
func maxSubArray(nums []int) int {
	result, tmp := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if tmp < 0 {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}

		if tmp > result {
			result = tmp
		}
	}
	return result
}

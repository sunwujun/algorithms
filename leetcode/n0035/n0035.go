package n0035

// [35] 搜索插入位置

// 暴力解法
func searchInsert1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// 二分法
func searchInsert(nums []int, target int) int {
	left, right, middle := 0, len(nums), 0
	for left < right {
		middle = (left + right) / 2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

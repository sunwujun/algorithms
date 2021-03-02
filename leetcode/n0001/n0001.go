package n0001

// [1] 两数之和

// 暴力解法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 借助 map
func twoSum2(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if index, ok := hash[target-nums[i]]; ok {
			return []int{index, i}
		}
		hash[nums[i]] = i
	}
	return nil
}

package n0016

import (
	"math"
	"sort"
)

// [16] 最接近的三数之和

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var n, closest = len(nums), math.MaxInt32
	for i := 0; i < n-2; i++ {
		for left, right := i+1, n-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}
		}
	}
	return closest
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

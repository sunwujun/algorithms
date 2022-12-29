package n0350

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	m := map[int]int{}
	for _, n := range nums1 {
		m[n] += 1
	}

	for _, n := range nums2 {
		if m[n] >= 1 {
			res = append(res, n)
			m[n] -= 1
		}
	}

	return res
}

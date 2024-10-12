func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	seen := make(map[int]bool)
	for _, num := range nums1 {
		seen[num] = true
	}

	for _, num := range nums2 {
		if seen[num] {
			result = append(result, num)
			seen[num] = false
		}
	}

	return result
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	res := 0

	for _, v := range nums {
		set[v] = true
	}

	for _, v := range nums {
		if !set[v-1] {
			length := 0

			for set[v+length] {
				length++
			}

			if length > res {
				res = length
			}
		}
	}

	return res
}
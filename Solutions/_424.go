func characterReplacement(s string, k int) int {
	left := 0
	right := 0
	maxCount := 0
	count := make([]int, 26)
	maxf := 0

	for right < len(s) {
		count[s[right]-'A']++
		maxf = max(maxf, count[s[right]-'A'])

		if (right-left+1)-maxf > k {
			count[s[left]-'A']--
			left++
		}

		maxCount = max(maxCount, right-left+1)
		right++
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

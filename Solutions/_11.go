func maxArea(height []int) (maxWater int) {
	left := 0
	right := len(height) - 1

	for left < right {
		currentWater := min(height[left], height[right]) * (right - left)
		maxWater = max(maxWater, currentWater)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

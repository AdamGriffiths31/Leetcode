func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		current := numbers[left] + numbers[right]
		if current > target {
			right--
		} else if current < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}

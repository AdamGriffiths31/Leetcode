func removeElement(nums []int, val int) int {
	pointer := 0
	for k, x := range nums {
		if x != val {
			nums[pointer] = nums[k]
			pointer++
		}
	}
	return pointer
}

func increasingTriplet(nums []int) bool {
	i := math.MaxInt
	j := math.MaxInt

	for _, num := range nums {
		if num <= i {
			i = num
		} else if num <= j {
			j = num
		} else {
			return true
		}
	}

	return false
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	result := []bool{}
	for _, i := range candies {
		if i > max {
			max = i
		}
	}

	for _, i := range candies {
		result = append(result, (i+extraCandies) >= max)
	}

	return result
}

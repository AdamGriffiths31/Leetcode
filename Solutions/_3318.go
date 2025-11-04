func findXSum(nums []int, k int, x int) []int {
	result := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		mapCount := make(map[int]int)

		for c := i; c < i+k; c++ {
			mapCount[nums[c]]++
		}

		pairs := []struct {
			value int
			freq  int
		}{}

		for val, count := range mapCount {
			pairs = append(pairs, struct {
				value int
				freq  int
			}{val, count})
		}

		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].freq != pairs[j].freq {
				return pairs[i].freq > pairs[j].freq
			}
			return pairs[i].value > pairs[j].value
		})

		sum := 0
		for i := 0; i < x && i < len(pairs); i++ {
			sum += pairs[i].value * pairs[i].freq
		}
		result = append(result, sum)
	}

	return result
}

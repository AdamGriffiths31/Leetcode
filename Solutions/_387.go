func firstUniqChar(s string) int {

	countMap := map[rune]int{}

	for _, value := range s {
		countMap[value] += 1
	}

	for i, value := range s {
		if countMap[value] == 1 {
			return i
		}
	}

	return -1
}

func uniqueOccurrences(arr []int) bool {
	hashmap := make(map[int]int)
	count := make(map[int]int)

	for _, val := range arr {
		hashmap[val]++
	}

	for _, val := range hashmap {
		if _, ok := count[val]; ok {
			return false
		}
		count[val]++
	}
	return true
}
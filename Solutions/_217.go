func containsDuplicate(nums []int) bool {
	hashmap := make(map[int]int)
	for i, v := range nums {
		if _, ok := hashmap[v]; ok {
			return true
		}
		hashmap[v] = i
	}
	return false
}

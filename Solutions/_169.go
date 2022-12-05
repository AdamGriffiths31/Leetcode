func majorityElement(nums []int) int {
	hashmap := make(map[int]int)
	for _, v := range nums {
		hashmap[v]++
	}
	for num, freq := range hashmap {
		if (freq) > len(nums)/2 {
			return num
		}
	}
	return 0
}
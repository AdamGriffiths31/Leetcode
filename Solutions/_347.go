func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	freq := map[int][]int{}

	for _, v := range nums {
		hashmap[v]++
	}

	for n, c := range hashmap {
		freq[c] = append(freq[c], n)
	}

	var res []int
	for i := len(hashmap) + 1; i >= 0; i-- {
		res = append(res, freq[i]...)
		if len(res) == k {
			return res
		}
	}
	return res
}
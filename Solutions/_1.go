package main

func twoSum(nums []int, target int) []int {
	var returnVal []int
	hashmap := make(map[int]int)
	for i, v := range nums {
		b := target - v
		if k, ok := hashmap[b]; ok {
			returnVal = append(returnVal, k, i)
			return returnVal
		}
		hashmap[v] = i
	}
	return returnVal
}

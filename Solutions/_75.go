func sortColors(nums []int) {
	bucket := make(map[int]int)
	for _, val := range nums {
		bucket[val]++
	}
	pointer := 0
	for i := 0; i <= 2; i++ {
		bucketCount := bucket[i]
		for x := bucketCount; x != 0; x-- {
			nums[pointer] = i
			pointer++
		}
	}
	fmt.Println(nums)
}

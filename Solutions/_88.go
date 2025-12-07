package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	arrayOneHighest := m - 1
	arrayTwoHighest := n - 1
	current := len(nums1) - 1

	for arrayOneHighest >= 0 && arrayTwoHighest >= 0 {
		if nums1[arrayOneHighest] >= nums2[arrayTwoHighest] {
			nums1[current] = nums1[arrayOneHighest]
			arrayOneHighest--
		} else {
			nums1[current] = nums2[arrayTwoHighest]
			arrayTwoHighest--
		}
		current--
	}

	for arrayTwoHighest >= 0 {
		nums1[current] = nums2[arrayTwoHighest]
		arrayTwoHighest--
		current--
	}

	fmt.Println(nums1)
}

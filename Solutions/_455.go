package main

import "fmt"

func main() {
	res := findContentChildren([]int{1, 2, 3}, []int{1, 1})
	fmt.Printf("res: %v\n", res)
	res = findContentChildren([]int{1, 2}, []int{1, 2, 3})
	fmt.Printf("res: %v\n", res)
}

func findContentChildren(g []int, s []int) int {
	g = sort(g)
	s = sort(s)
	res := 0
	for cookieIdx := 0; cookieIdx < len(s) && res < len(g); cookieIdx++ {
		if s[cookieIdx] >= g[res] {
			res++
		}
	}
	return res
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]

			}
		}
	}
	return arr
}

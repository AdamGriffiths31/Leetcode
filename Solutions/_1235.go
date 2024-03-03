package main

import (
	"fmt"
	"sort"
)

type job struct {
	startTime int
	endTime   int
	profit    int
}

var intervals []job
var cache = make(map[int]int)

func main() {

	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}

	fmt.Println(jobScheduling(startTime, endTime, profit))
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	cache = make(map[int]int)
	intervals = make([]job, len(startTime))
	for i := 0; i < len(startTime); i++ {
		intervals[i] = job{startTime[i], endTime[i], profit[i]}
	}
	sortFunc := func(i, j int) bool {
		return intervals[i].startTime < intervals[j].startTime
	}

	sort.Slice(intervals, sortFunc)
	return dfs(0)
}

func dfs(i int) int {
	if i == len(intervals) {
		return 0
	}

	if result, ok := cache[i]; ok {
		return result
	}

	result := dfs(i + 1)

	j := i + 1
	for j < len(intervals) {
		if intervals[j].startTime >= intervals[i].endTime {
			break
		}
		j++
	}

	result = max(result, intervals[i].profit+dfs(j))
	cache[i] = result
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

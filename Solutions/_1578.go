func minCost(colors string, neededTime []int) int {
	totalCost := 0

	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			totalCost += min(neededTime[i], neededTime[i+1])
			neededTime[i+1] = max(neededTime[i], neededTime[i+1])
		}
	}

	return totalCost
}

func findBall(grid [][]int) []int {
	res := make([]int, len(grid[0]))

	for i := 0; i < len(grid[0]); i++ {
		col := i
		for row := 0; row < len(grid); row++ {
			dir := grid[row][col]
			if col+dir >= 0 && col+dir < len(grid[0]) && dir == grid[row][col+dir] {
				col += dir
			} else {
				col = -1
				break
			}
		}
		res[i] = col
	}
	return res
}
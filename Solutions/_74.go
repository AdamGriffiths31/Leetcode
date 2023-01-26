func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	top := 0
	bottom := rows - 1

	for top <= bottom {
		row := (top + bottom) / 2
		if target > matrix[row][cols-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bottom = row - 1
		} else {
			break
		}
	}

	if !(top <= bottom) {
		return false
	}

	row := (top + bottom) / 2
	left := 0
	right := cols - 1
	for left <= right {
		middle := (left + right) / 2
		if target > matrix[row][middle] {
			left = middle + 1
		} else if target < matrix[row][middle] {
			right = middle - 1
		} else {
			return true
		}
	}
	return false
}

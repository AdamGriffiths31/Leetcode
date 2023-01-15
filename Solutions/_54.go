func spiralOrder(matrix [][]int) []int {
	var res []int
	left := 0
	right := len(matrix[0])
	top := 0
	bottom := len(matrix)

	for left < right && top < bottom {
		//top row
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		//right col
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right--

		if !(left < right && top < bottom) {
			break
		}

		//bottom row
		for i := right - 1; i != left-1; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom--

		//left row
		for i := bottom - 1; i != top-1; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
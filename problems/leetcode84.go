package problems
func dpMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	return matrix
}
func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	dpHeight, dpLeft, dpRight := dpMatrix(rows, cols), dpMatrix(rows, cols), dpMatrix(rows, cols)
	smallerInt := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	biggerInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	area := 0
	for i := 0; i < rows; i++ {
		curLeft := 0
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i > 0 {
					dpLeft[i][j] = biggerInt(curLeft, dpLeft[i-1][j])
					dpHeight[i][j] = dpHeight[i-1][j] + 1
				} else {
					dpLeft[i][j] = curLeft
					dpHeight[i][j] = 1
				}
			} else {
				curLeft = j + 1
				//FIXME remember setting default value
				dpLeft[i][j] = 0
				dpHeight[i][j] = 0
			}

		}
		curRight := cols
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if i > 0 {
					dpRight[i][j] = smallerInt(curRight, dpRight[i-1][j])

				} else {
					dpRight[i][j] = curRight
				}
			} else {
				curRight = j
				//FIXME remember setting default value
				dpRight[i][j] = cols
			}
			tmpArea := (dpRight[i][j] - dpLeft[i][j]) * dpHeight[i][j]
			if tmpArea > area {
				area = tmpArea
			}
		}

	}
	return area
}


package solution

func luckyNumbers(matrix [][]int) []int {
	var minRows []int
	var maxCols []int
	var intersection []int

	for i := 0; i < len(matrix); i++ {
		minRowVal := matrix[i][0]
		for j := 1; j < len(matrix[i]); j++ {
			if minRowVal > matrix[i][j] {
				minRowVal = matrix[i][j]
			}
		}
		minRows = append(minRows, minRowVal)
	}

	for j := 0; j < len(matrix[0]); j++ {
		maxColVal := matrix[0][j]
		for i := 1; i < len(matrix); i++ {
			if maxColVal < matrix[i][j] {
				maxColVal = matrix[i][j]
			}
		}
		maxCols = append(maxCols, maxColVal)
	}

	inArray := func(val int, array []int) bool {
		for i := 0; i < len(array); i++ {
			if array[i] == val {
				return true
			}
		}

		return false
	}

	for _, v := range minRows {
		isHave := true
		for i := 1; i < len(maxCols); i++ {
			if !inArray(v, maxCols) {
				isHave = false
				break
			}
		}

		if isHave {
			intersection = append(intersection, v)
		}
	}

	if len(intersection) == len(matrix) {
		maxVal := intersection[0]

		for i := 1; i < len(intersection); i++ {
			maxVal = max(maxVal, intersection[i])
		}

		return []int{maxVal}
	}

	return intersection
}

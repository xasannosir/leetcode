package solution

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	var result [][]int
	index := 0
	for i := 0; i < m; i++ {
		var arr []int
		for j := 0; j < n; j++ {
			arr = append(arr, original[index])
			index++
		}
		result = append(result, arr)
	}

	return result
}

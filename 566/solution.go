package solution

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])

	if n*m != r*c {
		return mat
	}

	arr := make([]int, n*m)

	for i, k := 0, 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[k] = mat[i][j]
			k++
		}
	}

	res := make([][]int, r)
	for i, k := 0, 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			res[i][j] = arr[k]
			k++
		}
	}

	return res
}

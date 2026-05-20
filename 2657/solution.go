package solution

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		pairs := make(map[int]int)

		for j := 0; j <= i; j++ {
			pairs[A[j]]++
		}

		for j := 0; j <= i; j++ {
			pairs[B[j]]++
		}

		count := 0
		for _, v := range pairs {
			if v > 1 {
				count++
			}
		}

		res[i] = count
	}

	return res
}

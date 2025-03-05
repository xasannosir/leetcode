package solution

func sortTheStudents(score [][]int, k int) [][]int {
	var numbers []int
	var indices []int
	var result [][]int

	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			if j == k {
				numbers = append(numbers, score[i][j])
				indices = append(indices, i)
				break
			}
		}
	}

	for i := 0; i < len(numbers)-1; i++ {
		flag := true
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				indices[j], indices[j+1] = indices[j+1], indices[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}

	for _, ind := range indices {
		result = append(result, score[ind])
	}

	return result
}

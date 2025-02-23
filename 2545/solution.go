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

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
				indices[i], indices[j] = indices[j], indices[i]
			}
		}
	}

	for _, ind := range indices {
		result = append(result, score[ind])
	}

	return result
}

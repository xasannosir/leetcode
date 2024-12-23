package solution

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var result []int
	var others []int

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr1[j] == arr2[i] {
				result = append(result, arr1[j])
			}
		}
	}

	for i := 0; i < len(arr1); i++ {
		isHave := false
		for j := 0; j < len(result); j++ {
			if arr1[i] == result[j] {
				isHave = true
				break
			}
		}
		if !isHave {
			others = append(others, arr1[i])
		}
	}

	for i := 0; i < len(others); i++ {
		for j := i + 1; j < len(others); j++ {
			if others[i] > others[j] {
				temp := others[i]
				others[i] = others[j]
				others[j] = temp
			}
		}
	}

	for i := 0; i < len(others); i++ {
		result = append(result, others[i])
	}

	return result
}

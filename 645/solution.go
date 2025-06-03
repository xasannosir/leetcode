package solution

func findErrorNums(nums []int) []int {
	counter := make(map[int]int)
	checker := make(map[int]bool)
	res := make([]int, 2)

	for _, num := range nums {
		counter[num]++
	}

	for i := 1; i <= len(nums); i++ {
		checker[i] = false
	}

	for k, v := range counter {
		if v == 2 {
			res[0] = k
		}
		delete(counter, k)
		delete(checker, k)
	}

	for k, _ := range checker {
		res[1] = k
		delete(checker, k)
		break
	}

	return res
}

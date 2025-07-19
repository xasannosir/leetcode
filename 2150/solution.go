package solution

func findLonely(nums []int) []int {
	counter := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums {
		counter[num]++
	}

	for k, v := range counter {
		_, ok1 := counter[k+1]
		_, ok2 := counter[k-1]

		if v == 1 && !ok1 && !ok2 {
			result = append(result, k)
			delete(counter, k)
		}
	}

	return result
}

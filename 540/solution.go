package solution

func singleNonDuplicate(nums []int) int {
	counter := make(map[int]int)

	for _, val := range nums {
		counter[val]++
	}

	for val, count := range counter {
		if count == 1 {
			return val
		}
	}

	return -1
}

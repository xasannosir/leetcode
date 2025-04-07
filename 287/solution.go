package solution

func findDuplicate(nums []int) int {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	for num, count := range counter {
		if count > 1 {
			return num
		}
	}

	return -1
}

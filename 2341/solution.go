package solution

func numberOfPairs(nums []int) []int {
	count := 0
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	for k, v := range counter {
		if v%2 == 0 {
			delete(counter, k)
			count += v / 2
		} else {
			counter[k] = 1
			count += v / 2
		}
	}

	return []int{count, len(counter)}
}

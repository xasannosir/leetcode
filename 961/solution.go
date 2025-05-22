package solution

func repeatedNTimes(nums []int) int {
	counter := make(map[int]int)
	n := len(nums)

	for _, val := range nums {
		counter[val]++
	}

	for k, v := range counter {
		if v == n/2 {
			return k
		}
	}

	return -1
}

package solution

func countBadPairs(nums []int) int64 {
	counter := make(map[int]int)

	for i, v := range nums {
		counter[v-i]++
	}

	sum := 0

	for _, v := range counter {
		sum += v
	}

	res := 0

	for _, v := range counter {
		res += v * (sum - v)
	}

	return int64(res / 2)
}

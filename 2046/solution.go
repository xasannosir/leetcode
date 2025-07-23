package solution

func isPossibleToSplit(nums []int) bool {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
		if counter[num] > 2 {
			return false
		}
	}

	return true
}

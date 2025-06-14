package solution

func minOperations(nums []int, k int) int {
	total := 0
	count := 0

	for _, num := range nums {
		total += num
	}

	for total >= 0 {
		if total%k == 0 {
			return count
		}

		count++
		total--
	}

	return count
}

package solution

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minInd := 0

		for j := 1; j < len(nums); j++ {
			if nums[j] < nums[minInd] {
				minInd = j
			}
		}

		nums[minInd] *= multiplier
	}

	return nums
}

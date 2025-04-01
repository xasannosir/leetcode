package solution

func maximumCount(nums []int) int {
	maximum, minimum := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			maximum++
		} else if nums[i] < 0 {
			minimum++
		}
	}

	return max(maximum, minimum)
}

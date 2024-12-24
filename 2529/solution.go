package solution

func maximumCount(nums []int) int {
	var maximum, minimum int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			maximum++
		} else if nums[i] < 0 {
			minimum++
		}
	}

	if maximum < minimum {
		return minimum
	} else {
		return maximum
	}
}

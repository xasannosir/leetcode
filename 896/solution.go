package solution

func isMonotonic(nums []int) bool {
	isIncrease := false
	isDecrease := false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			isIncrease = true
		} else if nums[i] > nums[i+1] {
			isDecrease = true
		} else {
			continue
		}

		if isIncrease && !isDecrease {
			continue
		} else if isDecrease && !isIncrease {
			continue
		} else {
			return false
		}
	}

	return true
}

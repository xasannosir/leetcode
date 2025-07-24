package solution

func findMiddleIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		leftSum, rightSum := 0, 0

		for j := 0; j < i; j++ {
			leftSum += nums[j]
		}

		for j := i + 1; j < len(nums); j++ {
			rightSum += nums[j]
		}

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

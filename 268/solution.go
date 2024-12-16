package solution

func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2

	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}

	return sum
}

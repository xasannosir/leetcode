package solution

func sumOfGoodNumbers(nums []int, k int) int {
	start, end, total := 0, len(nums)-1, 0

	for i := 0; i < len(nums); i++ {
		if i+k <= end && i-k >= start && nums[i] > nums[i+k] && nums[i] > nums[i-k] {
			total += nums[i]
		} else if i+k <= end && i-k < start && nums[i] > nums[i+k] {
			total += nums[i]
		} else if i-k >= start && i+k > end && nums[i] > nums[i-k] {
			total += nums[i]
		}
	}

	return total
}

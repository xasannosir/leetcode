package solution

func countPartitions(nums []int) int {
	count := 0
	n := len(nums)
	sum := 0
	left := 0
	right := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	right = sum

	for i := 0; i < n-1; i++ {
		left += nums[i]
		right -= nums[i]

		if (left-right)%2 == 0 {
			count++
		}
	}

	return count
}

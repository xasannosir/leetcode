package solution

func maxAdjacentDistance(nums []int) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}

		return n
	}

	n := len(nums)
	res := abs(nums[0] - nums[n-1])

	for i := 0; i < n-1; i++ {
		res = max(res, abs(nums[i]-nums[i+1]))
	}

	return res
}

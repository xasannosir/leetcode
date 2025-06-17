package solution

func minimumSum(nums []int) int {
	n := len(nums)
	res := -1

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] {
					if res == -1 {
						res = nums[i] + nums[j] + nums[k]
					} else {
						res = min(res, nums[i]+nums[j]+nums[k])
					}
				}
			}
		}
	}

	return res
}

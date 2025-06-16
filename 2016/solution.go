package solution

func maximumDifference(nums []int) int {
	res := -1

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				res = max(res, nums[j]-nums[i])
			}
		}
	}

	return res
}

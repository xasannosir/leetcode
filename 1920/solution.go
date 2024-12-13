package solution

func buildArray(nums []int) []int {
	var ans []int

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}

	return ans
}

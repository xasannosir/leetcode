package solution

func findMin(nums []int) int {
	res := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}

	return res
}

package solution

func shuffle(nums []int, n int) []int {
	result := make([]int, n*2)

	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		result[j] = nums[i]
		result[j+1] = nums[i+n]
	}

	return result
}

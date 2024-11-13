package solution

func getConcatenation(nums []int) []int {
	var result []int

	result = append(result, nums...)
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
	}

	return result
}

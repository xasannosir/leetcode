package solution

func runningSum(nums []int) []int {
	var (
		result []int
		sum    int
	)

	for _, num := range nums {
		sum += num
		result = append(result, sum)
	}

	return result
}

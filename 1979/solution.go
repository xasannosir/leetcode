package solution

func findGCD(nums []int) int {
	minVal := nums[0]
	maxVal := nums[0]

	for i := 1; i < len(nums); i++ {
		minVal = min(minVal, nums[i])
		maxVal = max(maxVal, nums[i])
	}

	greatest := 1

	for i := 2; i <= minVal; i++ {
		if minVal%i == 0 && maxVal%i == 0 {
			greatest = i
		}
	}

	return greatest
}

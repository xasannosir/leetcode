package solution

func findFinalValue(nums []int, original int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == original {
			original = original * 2
			i = -1
		}
	}

	return original
}

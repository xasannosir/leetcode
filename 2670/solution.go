package solution

func distinctDifferenceArray(nums []int) []int {
	diff := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		prefix := make(map[int]struct{})
		for j := 0; j <= i; j++ {
			prefix[nums[j]] = struct{}{}
		}

		suffix := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			suffix[nums[j]] = struct{}{}
		}

		diff = append(diff, len(prefix)-len(suffix))
	}

	return diff
}

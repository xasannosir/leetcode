package solution

func numOfPairs(nums []string, target string) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				count++
			}
		}
	}

	return count
}

package solution

func numIdenticalPairs(nums []int) int {
	var count int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}

	return count
}

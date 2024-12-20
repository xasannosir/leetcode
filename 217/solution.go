package solution

func containsDuplicate(nums []int) bool {
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	return len(counter) != len(nums)
}

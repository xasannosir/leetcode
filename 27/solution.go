package solution

func removeElement(nums []int, val int) int {
	var count, length = 0, len(nums)

	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

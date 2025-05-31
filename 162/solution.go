package solution

func findPeakElement(nums []int) int {
	count, index := 1, 0
	counter := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		} else if count != 0 {
			counter[count] = i
			count = 0
		}
	}

	if count != 0 {
		counter[count] = len(nums) - 1
	}

	for _, val := range counter {
		index = max(index, val)
	}

	return index
}

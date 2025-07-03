package solution

func arithmeticTriplets(nums []int, diff int) int {
	total := 0

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					total++
				}
			}
		}
	}

	return total
}

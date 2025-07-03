package solution

func countKDifference(nums []int, k int) int {
	abs := func(num int) int {
		if num > 0 {
			return num
		}

		return -num
	}

	total := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				total++
			}
		}
	}

	return total
}

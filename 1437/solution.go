package solution

func kLengthApart(nums []int, k int) bool {
	prev := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 && prev == -1 {
			prev = i
		} else if nums[i] == 1 && prev != -1 {
			if i-prev-1 < k {
				return false
			}

			prev = i
		}
	}

	return true
}

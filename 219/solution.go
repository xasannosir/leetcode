package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if nums[i] == nums[j] && abs(i-j) <= k {
				return true
			}
		}
	}

	return false
}

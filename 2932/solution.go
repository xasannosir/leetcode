package solution

func maximumStrongPairXor(nums []int) int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	choices := make([]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= min(nums[i], nums[j]) {
				choices = append(choices, nums[i]^nums[j])
			}
		}
	}

	var res int
	if len(choices) > 0 {
		res = choices[0]
	} else {
		return 0
	}

	for i := 1; i < len(choices); i++ {
		res = max(res, choices[i])
	}

	return res
}

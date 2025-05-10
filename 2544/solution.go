package solution

func alternateDigitSum(n int) int {
	nums := make([]int, 0)

	for n != 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	sum := 0
	sign := true

	for i := 0; i < len(nums); i++ {
		if sign {
			sum += nums[i]
		} else {
			sum += nums[i] * -1
		}

		sign = !sign
	}

	return sum
}

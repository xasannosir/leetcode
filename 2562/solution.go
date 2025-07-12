package solution

func findTheArrayConcVal(nums []int) int64 {
	var res int64

	build := func(a, b int) int64 {
		bStr := ""

		for b != 0 {
			bStr = string(rune(b%10+48)) + bStr
			b /= 10
		}

		for i := 0; i < len(bStr); i++ {
			a *= 10
			a += int(rune(bStr[i] - 48))
		}

		return int64(a)
	}

	for len(nums) != 0 {
		if len(nums) > 1 {
			res += build(nums[0], nums[len(nums)-1])
			nums = nums[1 : len(nums)-1]
		} else if len(nums) == 1 {
			res += int64(nums[0])
			break
		}
	}

	return res
}

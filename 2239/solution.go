package solution

func findClosestNumber(nums []int) int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	diff := abs(0 - nums[0])
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if diff > abs(0-nums[i]) {
			diff = abs(0 - nums[i])
			res = nums[i]
		} else if diff == abs(0-nums[i]) {
			res = max(res, nums[i])
		}
	}

	return res
}

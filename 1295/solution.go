package solution

func findNumbers(nums []int) int {
	var result int

	findNumberLength := func(num int) int {
		var count int

		for num > 0 {
			count++
			num = num / 10
		}

		return count
	}

	for i := 0; i < len(nums); i++ {
		if findNumberLength(nums[i])%2 == 0 {
			result++
		}
	}

	return result
}

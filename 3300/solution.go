package solution

func minElement(nums []int) int {
	var findSum func(num int) int

	findSum = func(num int) int {
		if num < 10 {
			return num
		}

		return num%10 + findSum(num/10)
	}

	minimum := findSum(nums[0])

	for i := 0; i < len(nums); i++ {
		minimum = min(minimum, findSum(nums[i]))
	}

	return minimum
}

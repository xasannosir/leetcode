package solution

func differenceOfSum(nums []int) int {
	digitSum, elemSum := 0, 0

	var findDigitSum func(num int) int

	findDigitSum = func(num int) int {
		if num > 9 {
			return num%10 + findDigitSum(num/10)
		}

		return num
	}

	abs := func(num int) int {
		if num < 0 {
			return num * -1
		}

		return num
	}

	for _, num := range nums {
		elemSum += num
		digitSum += findDigitSum(num)
	}

	return abs(elemSum - digitSum)
}

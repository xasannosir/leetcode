package solution

func findNumbers(nums []int) int {
	count := 0

	var countOfDigits func(num int, contain int) int

	countOfDigits = func(num int, contain int) int {
		if num == 0 {
			return contain
		}

		return countOfDigits(num/10, contain+1)
	}

	for _, num := range nums {
		if countOfDigits(num, 0)%2 == 0 {
			count++
		}
	}

	return count
}

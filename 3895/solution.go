package solution

func countDigitOccurrences(nums []int, digit int) int {
	count := 0

	for _, num := range nums {
		for num != 0 {
			if num%10 == digit {
				count++
			}

			num /= 10
		}
	}

	return count
}

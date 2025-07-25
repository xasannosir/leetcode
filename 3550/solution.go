package solution

func smallestIndex(nums []int) int {
	digitSum := func(num int) int {
		total := 0

		for num != 0 {
			total += num % 10
			num /= 10
		}

		return total
	}

	for i, num := range nums {
		if digitSum(num) == i {
			return i
		}
	}

	return -1
}

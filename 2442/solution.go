package solution

func countDistinctIntegers(nums []int) int {
	reverse := func(num int) int {
		res := 0

		for num != 0 {
			res = res*10 + (num % 10)
			num /= 10
		}

		return res
	}

	var twice []int
	copy(twice, nums)

	for _, num := range nums {
		twice = append(twice, num)
		twice = append(twice, reverse(num))
	}

	counter := make(map[int]int)

	for _, num := range twice {
		counter[num]++
	}

	return len(counter)
}

package solution

func mostFrequentEven(nums []int) int {
	counter := make(map[int]int)

	for _, num := range nums {
		if num%2 == 0 {
			counter[num]++
		}
	}

	res, resCount := -1, 0
	for num, count := range counter {
		if count > resCount {
			res = num
			resCount = count
		} else if count == resCount && num < res {
			res = num
		}
	}

	return res
}

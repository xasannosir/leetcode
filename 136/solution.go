package solution

func singleNumber(nums []int) int {
	var checker = make(map[int]int)

	for _, num := range nums {
		checker[num]++
	}

	for k, v := range checker {
		if v == 1 {
			return k
		}
	}

	return 0
}

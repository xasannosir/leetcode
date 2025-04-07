package solution

func singleNumber(nums []int) []int {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	var res []int
	for num, count := range counter {
		if count == 1 {
			res = append(res, num)
		}

		if len(res) == 2 {
			break
		}
	}

	return res
}

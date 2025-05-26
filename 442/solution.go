package solution

func findDuplicates(nums []int) []int {
	counter := make(map[int]int)
	res := make([]int, 0)

	for _, num := range nums {
		if counter[num] == 1 {
			res = append(res, num)
			delete(counter, num)
		} else {
			counter[num]++
		}
	}

	return res
}

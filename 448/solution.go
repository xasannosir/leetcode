package solution

func findDisappearedNumbers(nums []int) []int {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	var res []int
	for i := 1; i <= len(nums); i++ {
		if _, ok := counter[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}

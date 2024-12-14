package solution

func getSneakyNumbers(nums []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, num := range nums {
		counter[num]++
		if counter[num] == 2 {
			result = append(result, num)
		}
	}

	return result
}

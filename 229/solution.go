package solution

func majorityElement(nums []int) []int {
	count := len(nums) / 3
	counter := make(map[int]int)
	var result []int

	for _, val := range nums {
		counter[val]++
	}

	for k, v := range counter {
		if v > count {
			result = append(result, k)
		}
	}

	return result
}

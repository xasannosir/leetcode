package solution

func duplicateNumbersXOR(nums []int) int {
	counter := make(map[int]int)
	total := 0

	for _, num := range nums {
		counter[num]++

		if counter[num] == 2 {
			total ^= num
			delete(counter, num)
		}
	}

	return total
}

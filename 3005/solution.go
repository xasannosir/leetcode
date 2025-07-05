package solution

func maxFrequencyElements(nums []int) int {
	counter := make(map[int]int)
	maxFreq := 0
	total := 0

	for _, num := range nums {
		counter[num]++
	}

	for _, v := range counter {
		if v > maxFreq {
			maxFreq = v
		}
	}

	for _, v := range counter {
		if v == maxFreq {
			total += maxFreq
		}
	}

	return total
}

package solution

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	var numbers, counter []int

	index := 0
	for _, num := range nums {
		if ind, ok := freq[num]; ok {
			counter[ind]++
		} else {
			counter = append(counter, 1)
			numbers = append(numbers, num)
			freq[num] = index
			index++
		}
	}

	for i := 0; i < len(counter); i++ {
		for j := 0; j < len(counter)-i-1; j++ {
			if counter[j] < counter[j+1] {
				counter[j], counter[j+1] = counter[j+1], counter[j]
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers[:k]
}

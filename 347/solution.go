package solution

func topKFrequent(nums []int, k int) []int {
	numbers := make([]int, 0)
	counts := make([]int, 0)
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	for k, v := range counter {
		numbers = append(numbers, k)
		counts = append(counts, v)
	}

	for i := 0; i < len(counts); i++ {
		for j := 0; j < len(counts)-i-1; j++ {
			if counts[j] < counts[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				counts[j], counts[j+1] = counts[j+1], counts[j]
			}
		}
	}

	return numbers[:k]
}

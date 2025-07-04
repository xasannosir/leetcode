package solution

func frequencySort(nums []int) []int {
	counter := make(map[int]int)
	counts := make([]int, 0)
	numbers := make([]int, 0)
	res := make([]int, 0)

	for _, num := range nums {
		counter[num]++
	}

	for k, v := range counter {
		numbers = append(numbers, k)
		counts = append(counts, v)
	}

	for i := 0; i < len(counts); i++ {
		for j := 0; j < len(counts)-i-1; j++ {
			if counts[j] > counts[j+1] {
				counts[j], counts[j+1] = counts[j+1], counts[j]
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			} else if counts[j] == counts[j+1] && numbers[j] < numbers[j+1] {
				counts[j], counts[j+1] = counts[j+1], counts[j]
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	for i, val := range counts {
		for j := 0; j < val; j++ {
			res = append(res, numbers[i])
		}
	}

	return res
}

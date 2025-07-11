package solution

func minimumAverage(nums []int) float64 {
	n := len(nums)
	averages := make([]float64, 0)
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	for i := 0; i < n/2; i++ {
		minElement := 51
		maxElement := 0

		for k := range counter {
			minElement = min(minElement, k)
			maxElement = max(maxElement, k)
		}

		if counter[minElement] == 1 {
			delete(counter, minElement)
		} else {
			counter[minElement]--
		}

		if counter[maxElement] == 1 {
			delete(counter, maxElement)
		} else {
			counter[maxElement]--
		}

		averages = append(averages, (float64(minElement)+float64(maxElement))/2)
	}

	res := float64(averages[0])

	for i := 1; i < len(averages); i++ {
		res = min(res, averages[i])
	}

	return res
}

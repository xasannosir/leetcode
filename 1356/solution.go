package solution

func sortByBits(arr []int) []int {
	bins := make([]int, 0)

	binCount := func(num int) int {
		count := 0

		for num != 0 {
			count += num % 2
			num /= 2
		}

		return count
	}

	for _, num := range arr {
		bins = append(bins, binCount(num))
	}

	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if bins[j] > bins[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
				bins[j], bins[j+1] = bins[j+1], bins[j]
			} else if bins[j] == bins[j+1] {
				if arr[j] > arr[j+1] {
					flag = true
					arr[j], arr[j+1] = arr[j+1], arr[j]
					bins[j], bins[j+1] = bins[j+1], bins[j]
				}
			}
		}

		if !flag {
			break
		}
	}

	return arr
}

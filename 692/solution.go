package solution

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	res := make([]string, 0)
	keys := make([]string, 0)
	values := make([]int, 0)

	for _, word := range words {
		freq[word]++
	}

	for word, count := range freq {
		keys = append(keys, word)
		values = append(values, count)
	}

	sort := func(arr []int, arr2 []string) {
		for i := 0; i < len(arr); i++ {
			flag := false
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					arr2[j], arr2[j+1] = arr2[j+1], arr2[j]
					flag = true
				} else if arr[j] == arr[j+1] {
					if arr2[j] > arr2[j+1] {
						arr[j], arr[j+1] = arr[j+1], arr[j]
						arr2[j], arr2[j+1] = arr2[j+1], arr2[j]
						flag = true
					}
				}
			}

			if !flag {
				break
			}
		}
	}

	sort(values, keys)

	for i := 0; i < k; i++ {
		res = append(res, keys[i])
	}

	return res
}

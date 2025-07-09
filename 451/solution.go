package solution

func frequencySort(s string) string {
	counter := make(map[byte]int)
	values := make([]byte, 0)
	freq := make([]int, 0)
	res := ""

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	for k, v := range counter {
		values = append(values, k)
		freq = append(freq, v)
	}

	for i := 0; i < len(freq); i++ {
		for j := 0; j < len(freq)-i-1; j++ {
			if freq[j] < freq[j+1] {
				freq[j], freq[j+1] = freq[j+1], freq[j]
				values[j], values[j+1] = values[j+1], values[j]
			} else if freq[j] == freq[j+1] && values[j] < values[j+1] {
				freq[j], freq[j+1] = freq[j+1], freq[j]
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}

	for i := 0; i < len(freq); i++ {
		for j := 0; j < freq[i]; j++ {
			res += string(values[i])
		}
	}

	return res
}

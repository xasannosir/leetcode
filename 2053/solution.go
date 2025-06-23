package solution

func kthDistinct(arr []string, k int) string {
	counter := make(map[string]int)

	for _, s := range arr {
		counter[s]++
	}

	for _, s := range arr {
		if counter[s] == 1 {
			if k == 1 {
				return s
			}
			k--
		}
	}

	return ""
}

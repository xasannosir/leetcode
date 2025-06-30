package solution

func countWords(words1 []string, words2 []string) int {
	counter1 := make(map[string]int)
	counter2 := make(map[string]int)

	for _, word := range words1 {
		counter1[word]++
	}

	for _, word := range words2 {
		counter2[word]++
	}

	count := 0

	for k, v := range counter1 {
		if value, ok := counter2[k]; ok && v == 1 && value == 1 {
			count++
		}
	}

	return count
}

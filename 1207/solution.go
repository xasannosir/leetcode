package solution

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	occurences := make(map[int]int)

	for _, num := range arr {
		counter[num]++
	}

	for _, val := range counter {
		occurences[val]++
	}

	return len(occurences) == len(counter)
}

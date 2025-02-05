package solution

func checkIfPangram(sentence string) bool {
	counter := make(map[rune]bool)

	for _, val := range sentence {
		counter[val] = true
	}

	return len(counter) == 26
}

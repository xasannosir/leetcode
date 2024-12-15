package solution

func countConsistentStrings(allowed string, words []string) int {
	var count int

	for i := 0; i < len(words); i++ {
		var checker int
		for j := 0; j < len(words[i]); j++ {
			for k := 0; k < len(allowed); k++ {
				if words[i][j] == allowed[k] {
					checker++
					break
				}
			}
		}
		if checker == len(words[i]) {
			count++
		}
	}

	return count
}

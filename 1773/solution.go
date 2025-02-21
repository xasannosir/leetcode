package solution

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	index := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}

	count := 0
	for _, item := range items {
		if item[index[ruleKey]] == ruleValue {
			count++
		}
	}

	return count
}

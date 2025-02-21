package solution

func detectCapitalUse(word string) bool {
	var (
		allUpper, allLower, firstUpper []byte
	)

	for i := 0; i < len(word); i++ {
		if 65 <= word[i] && word[i] <= 90 {
			allUpper = append(allUpper, word[i])
			allLower = append(allLower, word[i]+32)
			firstUpper = append(firstUpper, word[i]+32)
		} else if 97 <= word[i] && word[i] <= 122 {
			allUpper = append(allUpper, word[i]-32)
			allLower = append(allLower, word[i])
			firstUpper = append(firstUpper, word[i])
		}
	}

	firstUpper[0] = firstUpper[0] - 32

	return word == string(allUpper) || word == string(allLower) || word == string(firstUpper)
}

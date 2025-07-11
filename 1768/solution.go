package solution

func mergeAlternately(word1 string, word2 string) string {
	res := ""

	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		res += string(word1[i])
		res += string(word2[i])
	}

	res += string(word1[i:])
	res += string(word2[i:])

	return res
}

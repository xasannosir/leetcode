package solution

func reverseWords(s string) string {
	var (
		result string
		helper []string
		word   string
	)

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word = string(s[i]) + word
		} else {
			helper = append(helper, word)
			word = ""
		}
	}
	helper = append(helper, word)

	for i := 0; i < len(helper); i++ {
		result += helper[i] + " "
	}
	result = result[:len(result)-1]

	return result
}

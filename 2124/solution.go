package solution

func checkString(s string) bool {
	bStartIndex := -1
	aEndIndex := -1

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aEndIndex = i
		} else if bStartIndex == -1 {
			bStartIndex = i
		}
	}

	if bStartIndex == -1 || aEndIndex == -1 {
		return true
	}

	return aEndIndex < bStartIndex
}

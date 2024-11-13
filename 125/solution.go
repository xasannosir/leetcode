package solution

func isPalindrome(s string) bool {
	var soft string

	for i := 0; i < len(s); i++ {
		value := s[i]
		if 65 <= value && value <= 90 {
			soft += string(value + 32)
		} else if 97 <= value && value <= 122 {
			soft += string(value)
		} else if 48 <= value && value <= 57 {
			soft += string(value)
		}
	}

	for i, j := 0, len(soft)-1; i < j; i, j = i+1, j-1 {
		if soft[i] != soft[j] {
			return false
		}
	}

	return true
}

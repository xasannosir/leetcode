package solution

func defangIPaddr(address string) string {
	var result string

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			result += "[.]"
		} else {
			result += string(address[i])
		}
	}

	return result
}

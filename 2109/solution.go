package solution

func addSpaces(s string, spaces []int) string {
	res := make([]byte, 0)

	for i, index := 0, 0; i < len(s); i++ {
		if index < len(spaces) && i == spaces[index] {
			res = append(res, ' ')
			index++
		}
		res = append(res, s[i])
	}

	return string(res)
}

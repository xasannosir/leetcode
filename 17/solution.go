package solution

func letterCombinations(digits string) []string {
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}

	for _, digit := range digits {
		tmp := []string{}

		for _, letter := range m[int(digit)-48] {
			if len(res) == 0 {
				tmp = append(tmp, string(letter))
				continue
			}

			for _, prev := range res {
				tmp = append(tmp, string(prev)+string(letter))
			}
		}

		res = tmp
	}

	return res
}

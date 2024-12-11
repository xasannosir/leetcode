package solution

func validStrings(n int) []string {
	var arr []string
	var add func(text string, num int) []string

	add = func(text string, num int) []string {
		if len(text) == num {
			arr = append(arr, text)
			return arr
		}

		if len(text) == 0 {
			add("0", num)
			add("1", num)
		} else if text[len(text)-1] == '0' {
			add(text+"1", num)
		} else {
			add(text+"0", num)
			add(text+"1", num)
		}

		return arr
	}

	return add("", n)
}

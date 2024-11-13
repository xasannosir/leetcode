package solution

func scoreOfString(s string) int {
	var sum int64

	for i := 0; i < len(s)-1; i++ {
		value := int64(s[i]) - int64(s[i+1])

		if value < 0 {
			sum += value * -1
		} else {
			sum += value
		}
	}

	return int(sum)
}

package solution

func findTheDifference(s string, t string) byte {
	sumT := 0
	sumS := 0

	for i := 0; i < len(t); i++ {
		sumT += int(t[i])
	}

	for i := 0; i < len(s); i++ {
		sumS += int(s[i])
	}

	return byte(sumT - sumS)
}

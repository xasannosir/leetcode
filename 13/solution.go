package solution

func romanToInt(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var number = 0
	for i := 0; i < len(s)-1; i++ {
		if romans[string(s[i])] < romans[string(s[i+1])] {
			number -= romans[string(s[i])]
		} else {
			number += romans[string(s[i])]
		}
	}

	number += romans[string(s[len(s)-1])]

	return number
}

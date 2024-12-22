package solution

func areNumbersAscending(s string) bool {
	var numbers []int
	var number int

	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			number = number*10 + int(s[i]-48)
		} else {
			if number > 0 {
				numbers = append(numbers, number)
				number = 0
			}
		}
	}

	if number > 0 {
		numbers = append(numbers, number)
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			return false
		}
	}

	return true
}

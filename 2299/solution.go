package solution

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return false
		}
	}

	special_characters := "!@#$%^&*()-+"
	isPunct := func(s byte) bool {
		for i := 0; i < len(special_characters); i++ {
			if special_characters[i] == s {
				return true
			}
		}

		return false
	}

	lower, upper, digit, character := false, false, false, false

	for i := 0; i < len(password); i++ {
		if 65 <= password[i] && password[i] <= 90 {
			upper = true
		} else if 97 <= password[i] && password[i] <= 122 {
			lower = true
		} else if 48 <= password[i] && password[i] <= 57 {
			digit = true
		} else if isPunct(password[i]) {
			character = true
		}
	}

	return lower && upper && digit && character
}

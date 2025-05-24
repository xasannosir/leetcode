package solution

func isStrictlyPalindromic(n int) bool {
	isPalindromeBaseK := func(num int, k int) bool {
		converted := ""

		for num != 0 {
			converted = string(rune(num%k+48)) + converted
			num /= k
		}

		bytes := []byte(converted)

		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}

		return string(bytes) == converted
	}

	for i := 2; i <= n-2; i++ {
		if !isPalindromeBaseK(n, i) {
			return false
		}
	}

	return true
}

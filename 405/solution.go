package solution

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"
	var result string
	n := uint32(num)

	for n > 0 {
		temp := n % 16
		result = string(hexChars[temp]) + result
		n /= 16
	}

	return result
}

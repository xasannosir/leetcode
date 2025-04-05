package solution

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		for i := len(num1); i < len(num2); i++ {
			num1 = "0" + num1
		}
	} else if len(num2) < len(num1) {
		for i := len(num2); i < len(num1); i++ {
			num2 = "0" + num2
		}
	}

	carry := 0
	res := ""

	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - 48)
		n2 := int(num2[i] - 48)
		total := n1 + n2 + carry
		number := total % 10
		carry = total / 10

		res = string(rune(number+48)) + res
	}

	if carry != 0 {
		res = string(rune(carry+48)) + res
	}

	return res
}

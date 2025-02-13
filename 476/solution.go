package solution

func findComplement(num int) int {
	pow := func(a, b int) int {
		res := 1

		for i := 0; i < b; i++ {
			res *= a
		}

		return res
	}

	binary := ""

	for num > 0 {
		bit := string(rune(num%2) + 48)

		if bit == "0" {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
		num /= 2
	}

	exponent := 0

	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			num += pow(2, exponent)
		}
		exponent++
	}

	return num
}

package solution

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	pow := func(a, b int) int {
		res := 1

		for i := 0; i < b; i++ {
			res *= a
		}

		return res
	}

	binary := ""

	for n > 0 {
		bit := string(rune(n%2) + 48)

		if bit == "0" {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
		n /= 2
	}

	exponent := 0
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			n += pow(2, exponent)
		}
		exponent++
	}

	return n
}

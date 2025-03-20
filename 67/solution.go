package solution

func addBinary(a string, b string) string {
	add := func(bin string, carry int) (string, int) {
		if carry == 0 {
			bin = "0" + bin
		} else if carry == 1 {
			bin = "1" + bin
			carry = 0
		} else if carry == 2 {
			bin = "0" + bin
			carry = 1
		} else if carry == 3 {
			bin = "1" + bin
			carry = 1
		}

		return bin, carry
	}

	var res string
	carry := 0

	if len(a) < len(b) {
		for len(a) != len(b) {
			a = "0" + a
		}
	} else {
		for len(b) != len(a) {
			b = "0" + b
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		carry = carry + int(a[i]-48) + int(b[i]-48)

		res, carry = add(res, carry)
	}

	res, carry = add(res, carry)

	if res[0] == '0' {
		res = res[1:]
	}

	return res
}

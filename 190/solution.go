package solution

func reverseBits(num uint32) uint32 {
	toBin := func(n uint32) string {
		bin := ""

		for n > 1 {
			bin = string(rune(n%2+48)) + bin
			n /= 2
		}

		bin = string(rune(n%2+48)) + bin

		for len(bin) != 32 {
			bin = "0" + bin
		}

		return bin
	}

	reverse := func(s string) string {
		bytes := []byte(s)
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}

		return string(bytes)
	}

	pow := func(a, b int) uint32 {
		res := 1

		for i := 0; i < b; i++ {
			res *= a
		}

		return uint32(res)
	}

	toNum := func(n string) uint32 {
		var num uint32
		exponent := 0

		for i := len(n) - 1; i >= 0; i-- {
			num += uint32(n[i]-48) * pow(2, exponent)
			exponent++
		}

		return num
	}

	return toNum(reverse(toBin(num)))
}

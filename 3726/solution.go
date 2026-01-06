package solution

import "strconv"

func removeZeros(n int64) int64 {
	str := strconv.Itoa(int(n))
	var res int64

	for i := 0; i < len(str); i++ {
		if str[i] != '0' {
			res = res*10 + int64(str[i]-48)
		}
	}

	return res
}

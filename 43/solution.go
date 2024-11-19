package solution

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}

	res := make([]byte, len(num1)+len(num2), len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			res[i+j+1] += (num1[i] - 48) * (num2[j] - 48)
			res[i+j] += res[i+j+1] / 10
			res[i+j+1] %= 10
		}
	}

	if res[0] == 0 {
		res = res[1:]
	}

	for i := range res {
		res[i] += 48
	}

	return string(res)
}

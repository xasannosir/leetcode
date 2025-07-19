package solution

func findDifferentBinaryString(nums []string) string {
	num2bin := func(num int) string {
		result := ""

		for num > 1 {
			result = string(rune(num%2+48)) + result
			num /= 2
		}

		return string(rune(num+48)) + result
	}

	build := func(bin string, length int) string {
		for len(bin) != length {
			bin = "0" + bin
		}

		return bin
	}

	checker := make(map[string]struct{})

	for _, num := range nums {
		checker[num] = struct{}{}
	}

	start := 0
	length := len(nums[0])

	for {
		random := build(num2bin(start), length)
		if _, ok := checker[random]; !ok {
			return random
		}
		start++
	}
}

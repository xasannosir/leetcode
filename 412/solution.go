package solution

func fizzBuzz(n int) []string {
	var result []string

	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			var (
				reverse int
				str     string
			)

			j := i

			for j != 0 {
				reverse = reverse*10 + j%10
				j /= 10
			}

			for reverse != 0 {
				str += string(rune(reverse%10 + 48))
				reverse /= 10
			}

			result = append(result, str)
		}
	}

	return result
}

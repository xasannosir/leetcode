package solution

func shortestToChar(s string, c byte) []int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	distances := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] != c {
			left, right := i-1, i+1

			for {
				if left >= 0 && s[left] == c {
					distances[i] = abs(i - left)
					break
				} else if right < len(s) && s[right] == c {
					distances[i] = abs(i - right)
					break
				}

				left--
				right++
			}
		}
	}

	return distances
}

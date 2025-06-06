package solution

type HashMap struct {
	Key int
	Val int
}

func countAndSay(n int) string {
	res := "1"

	numToPairs := func(s string) []HashMap {
		var pairs []HashMap

		for i := 0; i < len(s); i++ {
			count := 0
			j := i

			for ; j < len(s); j++ {
				if s[i] != s[j] {
					break
				}
				count++
			}

			pairs = append(pairs, HashMap{Key: int(s[i] - 48), Val: count})
			i = j - 1
		}

		return pairs
	}

	pairsToNum := func(pairs []HashMap) string {
		var num string

		for _, pair := range pairs {
			num += string(rune(pair.Val+48)) + string(rune(pair.Key+48))
		}

		return num
	}

	for i := 0; i < n-1; i++ {
		pairs := numToPairs(res)
		res = pairsToNum(pairs)
	}

	return res
}

package solution

func evenOddBit(n int) []int {
	bin := ""
	even := 0
	odd := 0

	for n != 0 {
		bin = string((n%2)+48) + bin
		n /= 2
	}

	for i, j := 0, len(bin)-1; j >= 0; i, j = i+1, j-1 {
		if bin[j] == '1' && i%2 == 0 {
			even++
		} else if bin[j] == '1' && i%2 == 1 {
			odd++
		}
	}

	return []int{even, odd}
}

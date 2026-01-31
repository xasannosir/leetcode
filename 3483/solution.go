package solution

func totalNumbers(digits []int) int {
	count := 0
	freq := make(map[int]struct{})
	n := len(digits)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i != j && j != k && k != i {
					num := (digits[i] * 100) + (digits[j] * 10) + digits[k]

					if num/100 != 0 && num%2 == 0 {
						if _, ok := freq[num]; !ok {
							count++
							freq[num] = struct{}{}
						}
					}
				}
			}
		}
	}

	return count
}

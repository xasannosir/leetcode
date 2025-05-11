package solution

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			count := 0
			for j := i; j < len(arr); j++ {
				if arr[j]%2 == 1 {
					count++
				} else {
					break
				}
			}

			if count == 3 {
				return true
			}
		}
	}

	return false
}

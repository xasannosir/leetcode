package solution

func findRelativeRanks(score []int) []string {
	var replace func(num int) string

	result := make([]string, len(score))
	order := make(map[int]int)

	sort := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			flag := false
			for j := 0; j < len(arr)-1-i; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = true
				}
			}

			if !flag {
				break
			}
		}
	}

	replace = func(num int) string {
		if num > 9 {
			return replace(num/10) + string(rune(num%10+48))
		}

		return string(rune(num + 48))
	}

	for ind, val := range score {
		order[val] = ind
	}

	sort(score)

	for ind, val := range score {
		if ind == 0 {
			result[order[val]] = "Gold Medal"
		} else if ind == 1 {
			result[order[val]] = "Silver Medal"
		} else if ind == 2 {
			result[order[val]] = "Bronze Medal"
		} else {
			result[order[val]] = replace(ind + 1)
		}
	}

	return result
}

package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool
	var maximum int

	for i := 0; i < len(candies); i++ {
		if candies[i] > maximum {
			maximum = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maximum {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

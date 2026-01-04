package solution

func sumFourDivisors(nums []int) int {
	total := 0

	for _, num := range nums {
		count := 2
		countSum := num + 1
		i := 2

		for ; i <= num/2; i++ {
			if num%i == 0 {
				count++
				countSum += i
			}

			if count > 4 {
				break
			}
		}

		if count == 4 {
			total += countSum
		}

		if i*i == num {
			continue
		}
	}

	return total
}

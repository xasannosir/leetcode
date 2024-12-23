package solution

func diagonalPrime(nums [][]int) int {
	maximum := 0

	isPrime := func(num int) bool {
		if num <= 2 {
			if num == 2 {
				return true
			} else {
				return false
			}
		}

		if num%2 == 0 {
			return false
		}

		i := 2

		for ; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if (i == j || i+j == len(nums)-1) && isPrime(nums[i][j]) && nums[i][j] > maximum {
				maximum = nums[i][j]
			}
		}
	}

	return maximum
}

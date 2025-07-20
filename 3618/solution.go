package solution

func splitArray(nums []int) int64 {
	sqrt := func(num int) int {
		start := 1

		for {
			if start*start == num {
				return start
			} else if start*start > num {
				return start - 1
			}
			start++
		}
	}

	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}

		if num == 2 {
			return true
		}

		if num%2 == 0 {
			return false
		}

		for i := 3; i <= sqrt(num); i += 2 {
			if num%i == 0 {
				return false
			}
		}

		return true
	}

	sum := func(arr []int) int {
		total := 0

		for _, num := range arr {
			total += num
		}

		return total
	}

	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	a := make([]int, 0)
	b := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if isPrime(i) {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}

	return int64(abs(sum(a) - sum(b)))
}
